# k8s-micro [英文](README.md)

Go Micro on kubernetes  

## 总览

此项目主要演示了[Go-Micro](https://github.com/micro/) 微服务框架如何运行在Kubernetes集群上，并通过调用APIServer的方式来进行服务发现注册及配置管理

## 特性

- Go-Micro作为微服务框架
- Protobuf作为编码协议
- GRPC作为RPC框架
- 基于Kubernetes的服务发现与注册
- 基于Kubernetes ConfigMap的配置管理
- 多服务案例
- 云原生应用

## 开始
- [安装 Go Micro]()
- [安装 Protobuf]()
- [创建 Kubernetes的命名空间]()
- [创建 RBAC]()
- [RPC服务案例]()
- [Web服务案例]()
- [多服务(Server/Client)运行案例]()
- [Go-micro(RPC/Web) on Kubernetes]()
- [使用 ConfigMap 作为配置管理]()

### 安装 Go Micro

```
go get github.com/micro/go-micro/v2
go get github.com/micro/go-plugins/registry/kubernetes/v2
```

### 安装 Protobuf

```
brew install protobuf
go get github.com/micro/micro/v2/cmd/protoc-gen-micro@master
protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. proto/greeter.proto
```

### 创建kubernetes的命名空间

**所有服务都运行在此命名空间下**

#### 命名空间清单写法如下

```
apiVersion: v1
kind: Namespace
metadata:
  name: go-micro
  namespace: go-micro
```

#### 部署命名空间

```
kubectl apply -f k8s/namespace.yaml
```

#### 查看结果

```
kubectl get ns |grep micro
go-micro          Active   36d
```

### 创建RBAC 

**对serviceaccount绑定操作pod及操作configmap的权限 会挂载到每个pod到/var/run/secrets/kubernetes.io/serviceaccount下面**

#### 创建Role

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
```

#### 创建ServiceAccount

```
apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: go-micro
  name: micro-service
```

#### 创建绑定关系

```
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
```



### Go Micro(RPC) on Kubernetes

**go-micro-srv**


```
make build or docker pull liuyao/go-micro-srv

```

#### Writing a Go Micro Service
```
func main() {
	service := micro.NewService(
		micro.Name(DefaultServiceName),
		micro.Server(grpcs.NewServer(server.Address(DefaultServerPort), server.Name(DefaultServiceName))),
		micro.Client(grpcc.NewClient()),
		micro.Registry(kubernetes.NewRegistry()),
	)
	service.Init()

	_ = proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
```
#### Deployment

**Here’s an example k8s deployment for a micro service**

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

#### Deploy with kubectl

```
kubectl apply -f k8s/deployment.yaml
```

#### Deplpy 
```
kubectl apply -f k8s/service.yaml
```

### Web服务运行案例

```
cd go-micro-web
make build or docker pull liuyao/go-micro-web
```


```
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```

```
kubectl get pods -n go-micro -o wide
NAME                            READY   STATUS    RESTARTS   AGE   IP           NODE             NOMINATED NODE   READINESS GATES
go-micro-srv-77c947dd6d-2rcj2   1/1     Running   0          16h   10.1.0.112   docker-desktop   <none>           <none>
go-micro-srv-77c947dd6d-474t5   1/1     Running   0          16h   10.1.0.113   docker-desktop   <none>           <none>
go-micro-web-56b457b9f7-f7lds   1/1     Running   0          65s   10.1.0.114   docker-desktop   <none>           <none>
go-micro-web-56b457b9f7-hvpg9   1/1     Running   0          65s   10.1.0.115   docker-desktop   <none>           <none>
```

```
kubectl logs go-micro-web-56b457b9f7-f7lds -n go-micro
2020-05-28 08:21:14  level=info service=web Listening on [::]:9200
```


```
kubectl get svc -n go-micro -o wide
kubectl describe svc go-micro-web -n go-micro 
```

kubectl get pods -n go-micro
NAME                              READY   STATUS    RESTARTS   AGE
go-micro-client-64b999f5d-m9ffk   1/1     Running   0          20m
go-micro-client-64b999f5d-wntkz   1/1     Running   0          20m



kubectl describe pod go-micro-client-64b999f5d-wntkz -n go-micro
Name:         go-micro-client-64b999f5d-wntkz
Namespace:    go-micro
Priority:     0
Node:         k8s-node-1/192.168.0.108
Start Time:   Tue, 28 Apr 2020 01:41:28 -0400
Labels:       app=go-micro-client
              micro.mu/selector-go-micro-client=service
              micro.mu/type=service
              pod-template-hash=64b999f5d
Annotations:  cni.projectcalico.org/podIP: 10.100.109.133/32
              cni.projectcalico.org/podIPs: 10.100.109.133/32
              micro.mu/service-go-micro-client:
                {"name":"go-micro-client","version":"latest","metadata":null,"endpoints":[],"nodes":[{"id":"go-micro-client-d227891e-29b1-4d7f-81ad-d53ae1...
Status:       Running
IP:           10.100.109.133
IPs:
  IP:           10.100.109.133
Controlled By:  ReplicaSet/go-micro-client-64b999f5d
Containers:
  go-micro-client:
    Container ID:   docker://5fffe1010fc626303834111c7e9a98c7c8a9effb2a2a9d4676e91a8d2fc2182f
    Image:          liuyao/go-micro-client:kubernetes
    Image ID:       docker-pullable://liuyao/go-micro-client@sha256:52704576092baf1639acb7f7cc63ff1e7f7a58b859187a5f5b2506be89a503ae
    Port:           <none>
    Host Port:      <none>
    State:          Running
      Started:      Tue, 28 Apr 2020 01:57:37 -0400
    Ready:          True
    Restart Count:  0
    Environment:    <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from micro-services-token-j9gr8 (ro)
      
                      

kubectl logs go-micro-client-64b999f5d-wntkz -n go-micro
2020-04-28 05:57:37  level=info Starting [service] go-micro-client
2020-04-28 05:57:37  level=info Server [grpc] Listening on [::]:46701
2020-04-28 05:57:37  level=info Registry [kubernetes] Registering node: go-micro-client-d227891e-29b1-4d7f-81ad-d53ae11bb7f6
Server Hello Yao!
Server Hello Yao!
Server Hello Yao!
Server Hello Yao!
Server Hello Yao!
Server Hello Yao!
Server Hello Yao!
Server Hello Yao!
Server Hello Yao!
Server Hello Yao!
Server Hello Yao!
Server Hello Yao!
Server Hello Yao!


### 使用 ConfigMap 作为配置管理
#### 原理
此处有图
https://10.96.0.1:443/api/v1/namespaces/go-micro/configmaps/go-micro-config"
#### 写一个configmaps


#### 编写代码

#### 运行

#### 查看

[root@k8s-master-1 k8s]# kubectl logs go-micro-config -n go-micro
map[DB_HOST:map[192.168.0.1:] DB_NAME:map[MICRO:] 
go:map[micro:map[srv:map[port:map[9100:map[tcp:map[addr:10.96.196.160 port:9100 proto:tcp]]] 
service:map[host:10.96.196.160 port:map[go:map[micro:map[srv:9100]]]]] 
web:map[port:map[9200:map[tcp:map[addr:10.96.218.32 port:9200 proto:tcp]]] 
service:map[host:10.96.218.32 port:map[go:map[micro:map[web:9200]]]]]]] 
home:/root hostname:go-micro-config kubernetes:map[port:map[443:map[tcp:tcp://10.96.0.1:443]] 
service:map[host:10.96.0.1 port:443]] path:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin]
