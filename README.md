# k8s-micro

go micro on kubernetes 

## Overview

## Features 

## Getting Started

- [Gin on Kubernetes Demo]()
- [Go-micro(Web) on Kubernetes Demo]()
- [Go-micro(RPC) on Kubernetes Demo]()

### Gin on Kubernetes Demo
    
```
cd go-gin-demo
make build or docker pull liuyao/go-gin-demo
kubectl apply -f k8s-PersistentVolumeClaim.yaml 
kubectl apply -f k8s-Deployment.yaml
kubectl apply -f k8s-Service.yaml
```
### Go-micro on Kubernetes Demo
```
cd go-micro-demo
make build or docker pull liuyao/go-micro-demo
kubectl apply -f k8s-Role.yaml
kubectl apply -f k8s-RoleBinding.yaml
kubectl apply -f k8s-PersistentVolumeClaim.yaml 
kubectl apply -f k8s-Deployment.yaml
kubectl apply -f k8s-Service.yaml
```

go get github.com/micro/micro/v2/cmd/protoc-gen-micro@master
protoc --proto_path=$GOPATH/src:. --micro_out=. --go_out=. greeter.proto