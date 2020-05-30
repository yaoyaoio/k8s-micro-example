# k8s-micro [简体中文](README_CN.md)

Go Micro on kubernetes  

## Overview

ing...

## Features

ing...

## Getting Started

- [Dependencies]()
- [Installing Go Micro]()
- [Installing Protobuf]()
- [Create Kubernetes NameSpace](#create-kubernetes-namespace)
- [Go Micro(RPC) on Kubernetes]()
- [Go Micro(Web) on Kubernetes]()
- [Go Micro(RPC) MultiService on Kubernetes]()
- [Go-micro(RPC/Web) on Kubernetes]()
- [Using ConfigMap]()

### Gin on Kubernetes Demo
    
```
cd go-gin-demo
make build or docker pull liuyao/go-gin-demo
kubectl apply -f k8s-PersistentVolumeClaim.yaml 
kubectl apply -f k8s-Deployment.yaml
kubectl apply -f k8s-Service.yaml
```


### Installing Protobuf

```
go get github.com/micro/micro/v2/cmd/protoc-gen-micro@master
protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. greeter.proto
```


### Create Kubernetes NameSpace

1. 资源清单
    ```
    apiVersion: v1
    kind: Namespace
    metadata:
      name: go-micro
      namespace: go-micro
    ```
2. 创建NameSpace
    ```
    kubectl apply -f k8s/namespace.yaml
    ```
3. 查看创建结果
    ```
    kubectl get ns |grep micro
    ```
### Go-micro on Kubernetes

```
cd go-micro-srv
make build or docker pull liuyao/go-micro-srv
kubectl apply -f k8s/role.yaml
kubectl apply -f k8s/roleBinding.yaml
kubectl apply -f k8s/persistentVolumeClaim.yaml 
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```

### Go Micro Web on Kubernetes
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









