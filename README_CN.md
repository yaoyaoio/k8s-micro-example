# k8s-micro [简体中文](README_CN.md)

Go Micro on kubernetes  

## 总揽



## 特性

- 使用Protobuf作为编码协议
- 使用GRPC作为RPC框架
- 基于Kubernetes的服务发现与注册
- 基于Kubernetes ConfigMap的配置管理
- 多服务案例
- 云原生应用

## 开始
- [安装 Go Micro]()
- [安装 Protobuf]()
- [创建 Kubernetes的命名空间]()
- [创建 RBAC]()
- [RPC服务案例](#go-microrpc-on-kubernetes)
- [Web服务案例](#go-microweb-on-kubernetes)
- [多服务(Server/Client)运行案例](#go-microrpc-multiservice-on-kubernetes)
- [Go-micro(RPC/Web) on Kubernetes](#go-microrpcweb-on-kubernetes)
- [使用 ConfigMap 作为配置管理](#using-configmap)

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