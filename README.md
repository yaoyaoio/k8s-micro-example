# k8s-micro [简体中文](README_CN.md)

Go Micro on kubernetes  

## Overview

This project mainly demonstrates how the [Go-Micro](https://github.com/micro/) runs on the kubernetes cluster, 
and through the way of calling apiserver to carry out service discovery registration and configuration management

## Features

- Protobuf
- GRPC
- Kubernetes
- MultiService Example

## Getting Started

- [Installing Go Micro](#installing-go-micro)
- [Installing Protobuf and Writing Proto**](#installing-protobuf)
- [Create Kubernetes Namespace](#create-kubernetes-namespace)
- [Create RBAC](#create-rbac)
- [Kubernetes Discovery]()
- [Go Micro(RPC) on Kubernetes](#go-microrpc-on-kubernetes)
- [Go Micro(Web) on Kubernetes](#go-microweb-on-kubernetes)
- [Go Micro(RPC) MultiService on Kubernetes](#go-microrpc-multiservice-on-kubernetes)
- [Go-micro(RPC/Web) on Kubernetes](#go-microrpcweb-on-kubernetes)
- [Using ConfigMap](#using-configmap)

### Installing Go Micro

```
go get github.com/micro/go-micro/v2
go get github.com/micro/go-plugins/registry/kubernetes/v2
```

### Installing Protobuf

```
brew install protobuf
go get github.com/micro/micro/v2/cmd/protoc-gen-micro@master
protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. proto/greeter.proto
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
```

### Create RBAC






### Go Micro(RPC) on Kubernetes

**go-micro-srv**

#### Writing a Go Micro Service

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
Deploy with kubectl

```
kubectl apply -f k8s/deployment.yaml
```


```

kubectl apply -f k8s/service.yaml



make build or docker pull liuyao/go-micro-srv
kubectl apply -f k8s/role.yaml
kubectl apply -f k8s/roleBinding.yaml
kubectl apply -f k8s/persistentVolumeClaim.yaml 
kubectl apply -f k8s/deployment.yaml
```

### Go Micro(Web) on Kubernetes
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

### Go Micro(RPC) MultiService on Kubernetes






### Go-micro(RPC/Web) on Kubernetes




### Using ConfigMap
#### 原理
此处有图
https://10.96.0.1:443/api/v1/namespaces/go-micro/configmaps/go-micro-config"
#### 写一个configmaps

#### 编写代码

#### 运行

#### 查看

[root@k8s-master-1 k8s]# kubectl logs go-micro-config -n go-micro
map[DB_HOST:map[192.168.0.1:] DB_NAME:map[MICRO:] go:map[micro:map[srv:map[port:map[9100:map[tcp:map[addr:10.96.196.160 port:9100 proto:tcp]]] service:map[host:10.96.196.160 port:map[go:map[micro:map[srv:9100]]]]] web:map[port:map[9200:map[tcp:map[addr:10.96.218.32 port:9200 proto:tcp]]] service:map[host:10.96.218.32 port:map[go:map[micro:map[web:9200]]]]]]] home:/root hostname:go-micro-config kubernetes:map[port:map[443:map[tcp:tcp://10.96.0.1:443]] service:map[host:10.96.0.1 port:443]] path:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin]
