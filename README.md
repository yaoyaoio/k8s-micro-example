# k8s-micro [简体中文](README_CN.md)

Go Micro on kubernetes  

## Overview

This project mainly demonstrates how the [Go-Micro](https://github.com/micro/) runs on the kubernetes cluster, 
and through the way of calling apiserver to carry out service discovery registration and configuration management
## Features

- Go-Micro
- Protobuf
- GRPC
- Kubernetes Discovery
- Kubernetes ConfigMap
- MultiService Example
- Cloud Native

## Getting Started
- [Installing Go Micro](#installing-go-micro)
- [Installing Protobuf and Writing Proto](#installing-protobuf)
- [Create Kubernetes Namespace](#create-kubernetes-namespace)
- [Create RBAC](#create-rbac)
- [Kubernetes Discovery]()
- [Go Micro(RPC) on Kubernetes](#go-microrpc-on-kubernetes)
- [Go Micro(Web) on Kubernetes](#go-microweb-on-kubernetes)
- [Go Micro(RPC) MultiService on Kubernetes](#go-microrpc-multiservice-on-kubernetes)
- [Go-micro(RPC/Web) on Kubernetes](#go-microrpcweb-on-kubernetes)
- [Using ConfigMap](#using-configmap)
- [TODO:Health]()

### Installing Go Micro


```
go get github.com/micro/go-micro/v2@v2.3.0
go get github.com/micro/go-plugins/registry/kubernetes/v2@v2.3.0
go get github.com/micro/go-plugins/config/source/configmap/v2@v2.3.0
```

### Installing Protobuf & Writing Proto

#### Installing Protobuf
```
brew install protobuf
go get github.com/micro/micro/v2/cmd/protoc-gen-micro@master
```
#### Writing Proto



```
syntax = "proto3";

service Greeter {
    rpc Hello (HelloRequest) returns (HelloResponse) {
    }
}

message HelloRequest {
    string name = 1;
}

message HelloResponse {
    string greeting = 2;
}
```

#### Generate
```
protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. proto/greeter.proto
ls proto
greeter.pb.micro.go greeter.proto greeter.proto
```

### Create Kubernetes Namespace

#### Writing a namespace
```
apiVersion: v1
kind: Namespace
metadata:
  name: go-micro
  namespace: go-micro
```
#### Deploying a NameSpace
```
kubectl apply -f k8s/namespace.yaml
```
#### Select Result
```
kubectl get ns |grep micro
go-micro          Active   36d
```

### Create RBAC


#### Writing Pod RBAC
```
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: micro-registry
  namespace: go-micro
rules:
  - apiGroups:
      - ""
    resources:
      - pods
    verbs:
      - get
      - list
      - patch
      - watch
---
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: micro-registry
  namespace: go-micro
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: micro-registry
subjects:
  - kind: ServiceAccount
    name: micro-services
    namespace: go-micro
---
apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: go-micro
  name: micro-service
```
#### Writing ConfigMap RBAC
```
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: micro-config
  namespace: go-micro
  labels:
    app: go-micro-config
rules:
  - apiGroups: [""]
    resources: ["configmaps"]
    verbs: ["get", "update", "list", "watch"]
---
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: micro-config
  namespace: go-micro
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: Role
  name: micro-config
subjects:
  - kind: ServiceAccount
    name: micro-services
    namespace: go-micro
---
apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: go-micro
  name: micro-services
```
#### Create
```
kubectl apply -f k8s/configmap-rbac.yaml
kubectl apply -f k8s/pod-rbac.yaml
```

### Kubernetes Discovery




ing...

### Go Micro(RPC) on Kubernetes

#### Deploy
```
kubectl apply -f go-micro-srv/k8s/deployment.yaml
kubectl apply -f go-micro-srv/k8s/service.yaml
```
#### Writing Code
```
import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	grpcc "github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-micro/v2/server"
	grpcs "github.com/micro/go-micro/v2/server/grpc"
	"github.com/micro/go-plugins/registry/kubernetes/v2"
	proto "github.com/yaoliu/k8s-micro/proto"
	_ "net/http/pprof"
)

type Greeter struct{}

func (g Greeter) Hello(ctx context.Context, request *proto.HelloRequest, response *proto.HelloResponse) error {
	response.Greeting = "Hello " + request.Name + "!"
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name(DefaultServiceName),
		micro.Server(grpcs.NewServer(server.Address(DefaultServerPort), server.Name(DefaultServiceName))),
		micro.Client(grpcc.NewClient()), 
		micro.Registry(kubernetes.NewRegistry()),//注册到Kubernetes
	)
	service.Init()

	_ = proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
```
#### Writing Dockerfile
```
FROM alpine

MAINTAINER liuyao@163.com

ADD ./server /server

EXPOSE 9100

CMD ["/server"]
```
#### Compile & Push Image
```
CGO_ENABLED=0 GOOS=linux go build -o server main.go
docker build -t liuyao/go-micro-srv:kubernetes .
docker push liuyao/go-micro-srv:kubernetes
```
#### Writing Deployment
```
apiVersion: apps/v1
kind: Deployment
metadata:
  namespace: go-micro
  name: go-micro-srv
spec:
  selector:
    matchLabels:
      app: go-micro-srv
  replicas: 2
  template:
    metadata:
      labels:
        app: go-micro-srv
    spec:
      containers:
        - name: go-micro-srv
          image: liuyao/go-micro-srv:kubernetes
          imagePullPolicy: Always
          ports:
            - containerPort: 9100
              name: rpc-port
```
#### Writing Service
```
apiVersion: v1
kind: Service
metadata:
  name: go-micro-srv
  namespace: go-micro
  labels:
    app: go-micro-srv
spec:
  ports:
    - port: 9100
      name: go-micro-srv
      targetPort: 9100
  selector:
    app: go-micro-srv
```
#### Deploy
```
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```
#### Select Run Status & Registry Status & Start Logs
```
kubectl get pods -n go-micro |grep go-micro-srv
go-micro-srv-6cc7848c6-4knrm      1/1     Running     0          17h
go-micro-srv-6cc7848c6-lf6wm      1/1     Running     0          17h
kubectl describe pod go-micro-srv-6cc7848c6-4knrm -n go-micro
···
Labels:       app=go-micro-srv
              micro.mu/selector-go-micro-srv=service
              micro.mu/type=service
              pod-template-hash=6cc7848c6
Annotations:  cni.projectcalico.org/podIP: 10.100.109.174/32
              cni.projectcalico.org/podIPs: 10.100.109.174/32
              micro.mu/service-go-micro-srv:
                {"name":"go-micro-srv","version":"latest","metadata":null,"endpoints":[{"name":"Greeter.Hello","request":{"name":"HelloRequest","type":"He.
···
查看日志
kubectl logs go-micro-srv-6cc7848c6-4knrm -n go-micro
2020-04-28 03:31:03  level=info Starting [service] go-micro-srv
2020-04-28 03:31:03  level=info Server [grpc] Listening on [::]:9100
2020-04-28 03:31:03  level=info Registry [kubernetes] Registering node: go-micro-srv-5c1ae799-d6be-48aa-b56c-be3457508bc5
```

### Go Micro(Web) on Kubernetes

#### Deploy
```
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```
#### Writing Code
```
import (
	"fmt"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/kubernetes/v2"
	"net/http"
	"os"
)

func main() {
	service := web.NewService(
		web.Name("go-micro-web"),
		web.Registry(kubernetes.NewRegistry()),
		web.Address(":9200"))

	service.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		podName := os.Getenv("HOSTNAME")
		_, _ = writer.Write([]byte(podName))
	})

	if err := service.Init(); err != nil {
		fmt.Println(err)
	}

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
```
#### Writing Dockerfile
```
FROM alpine

MAINTAINER liuyao@163.com

ADD ./web /web

EXPOSE 9200

CMD ["/web"]
```
#### Compile & Push Image
```
CGO_ENABLED=0 GOOS=linux go build -o web main.go
docker build -t liuyao/go-micro-web:kubernetes .
docker push liuyao/go-micro-web:kubernetes
``` 
#### Writing Deployment
```
apiVersion: apps/v1
kind: Deployment
metadata:
  namespace: go-micro
  name: go-micro-web
spec:
  selector:
    matchLabels:
      app: go-micro-web
  replicas: 2
  template:
    metadata:
      labels:
        app: go-micro-web
    spec:
      containers:
        - name: go-micro-web
          image: liuyao/go-micro-web:kubernetes
          imagePullPolicy: Always
          ports:
            - containerPort: 9200
              name: http-port
      serviceAccountName: micro-services
```
#### Writing Service
```
apiVersion: v1
kind: Service
metadata:
  name: go-micro-web
  namespace: go-micro
  labels:
    app: go-micro-web
spec:
  ports:
    - port: 9200
      name: go-micro-web
      targetPort: 9200
  selector:
    app: go-micro-web

```
#### Deploy
```
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```
#### Select Run Status & Registry Status & Start Logs
```
kubectl get pods -n go-micro -o wide
NAME                            READY   STATUS    RESTARTS   AGE   IP           NODE             NOMINATED NODE   READINESS GATES
go-micro-web-56b457b9f7-f7lds   1/1     Running   0          65s   10.1.0.114   docker-desktop   <none>           <none>
go-micro-web-56b457b9f7-hvpg9   1/1     Running   0          65s   10.1.0.115   docker-desktop   <none>           <none>
kubectl logs go-micro-web-56b457b9f7-f7lds -n go-micro
2020-05-28 08:21:14  level=info service=web Listening on [::]:9200
```


### Go Micro(RPC) MultiService on Kubernetes



#### Deploy
```
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```
#### Writing Code
```
import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	grpcc "github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-plugins/registry/kubernetes/v2"
	proto "github.com/yaoliu/k8s-micro/proto"
	"time"
)

var (
	DefaultServiceName = "go-micro-client"
	DefaultSrvName     = "go-micro-srv"
)

func main() {
	service := micro.NewService(
		micro.Name(DefaultServiceName),
		micro.Client(grpcc.NewClient()),
		micro.Registry(kubernetes.NewRegistry()),
	)
	service.Init()
	go func() {
		registryRPCHandler(service.Client())
	}()
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}

func registryRPCHandler(s client.Client) {
	timer := time.NewTicker(time.Second * 10)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			greeter := proto.NewGreeterService(DefaultSrvName, s)
			rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "Yao"})
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server", rsp.Greeting)
			}
		}
	}
}
```
#### Writing Dockerfile
```
FROM alpine

MAINTAINER liuyao@163.com

ADD ./client /client

EXPOSE 9200

CMD ["/client"]
```
#### Compile & Push Images
```
CGO_ENABLED=0 GOOS=linux go build  -o client main.go
docker build -t liuyao/go-micro-client:kubernetes .
docker push liuyao/go-micro-client:kubernetes
```
#### Writing Deployment
```
apiVersion: apps/v1
kind: Deployment
metadata:
  namespace: go-micro
  name: go-micro-client
spec:
  selector:
    matchLabels:
      app: go-micro-client
  replicas: 2
  template:
    metadata:
      labels:
        app: go-micro-client
    spec:
      containers:
        - name: go-micro-client
          image: liuyao/go-micro-client:kubernetes
          imagePullPolicy: Always
      serviceAccountName: micro-services
```
#### Depoly
```
kubectl apply -f k8s/deployment.yaml
```
#### Select Run Status & Registry Status & Start Logs
```
kubectl get pods -n go-micro
NAME                              READY   STATUS    RESTARTS   AGE
go-micro-client-64b999f5d-m9ffk   1/1     Running   0          20m
go-micro-client-64b999f5d-wntkz   1/1     Running   0          20m
kubectl describe pod go-micro-client-64b999f5d-wntkz -n go-micro
···
Labels:       app=go-micro-client
              micro.mu/selector-go-micro-client=service
              micro.mu/type=service
              pod-template-hash=64b999f5d
Annotations:  cni.projectcalico.org/podIP: 10.100.109.133/32
              cni.projectcalico.org/podIPs: 10.100.109.133/32
              micro.mu/service-go-micro-client:
                {"name":"go-micro-client","version":"latest","metadata":null,"endpoints":[],"nodes":[{"id":"go-micro-client-d227891e-29b1-4d7f-81ad-d53ae1...
···
kubectl logs go-micro-client-64b999f5d-wntkz -n go-micro
2020-04-28 05:57:37  level=info Starting [service] go-micro-client
2020-04-28 05:57:37  level=info Server [grpc] Listening on [::]:46701
2020-04-28 05:57:37  level=info Registry [kubernetes] Registering node: go-micro-client-d227891e-29b1-4d7f-81ad-d53ae11bb7f6
Server Hello Yao!
```
### Using ConfigMap

#### Writing Config Map
```
apiVersion: v1
kind: ConfigMap
metadata:
  name: go-micro-config
  namespace: go-micro
data:
  DB_NAME: MICRO
  DB_HOST: 192.168.0.1
```
#### Writing Code
```
import (
	"fmt"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/env"
	"github.com/micro/go-plugins/config/source/configmap/v2"
)

var (
	DefaultNamespace  = "go-micro"
	DefaultConfigName = "go-micro-config"
)

func main() {
	if cfg, err := config.NewConfig(); err == nil {
		err = cfg.Load(
			env.NewSource(),
			configmap.NewSource(
				configmap.WithName(DefaultConfigName),
				configmap.WithNamespace(DefaultNamespace)),
		)
		if err == nil {
			fmt.Println(cfg.Map())
		}
		fmt.Println(err)
	}
}
```
#### Writing Pod
```
apiVersion: v1
kind: Pod
metadata:
  name: go-micro-config
  namespace: go-micro
spec:
  containers:
    - name: go-micro-config
      image: liuyao/go-micro-config:kubernetes
      imagePullPolicy: Always
  restartPolicy: Never
  serviceAccountName: micro-services
```
#### Deploy
```
cd go-micro-config
kubectl apply -f k8s/pod.yaml
```
#### Select Result
```
[root@k8s-master-1 k8s]# kubectl logs go-micro-config -n go-micro
map[DB_HOST:map[192.168.0.1:] DB_NAME:map[MICRO:] 
go:map[micro:map[srv:map[port:map[9100:map[tcp:map[addr:10.96.196.160 port:9100 proto:tcp]]] 
service:map[host:10.96.196.160 port:map[go:map[micro:map[srv:9100]]]]] 
web:map[port:map[9200:map[tcp:map[addr:10.96.218.32 port:9200 proto:tcp]]] 
service:map[host:10.96.218.32 port:map[go:map[micro:map[web:9200]]]]]]] 
home:/root hostname:go-micro-config kubernetes:map[port:map[443:map[tcp:tcp://10.96.0.1:443]] 
service:map[host:10.96.0.1 port:443]] path:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin]
```