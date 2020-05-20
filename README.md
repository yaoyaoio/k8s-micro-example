# k8s-micro

用于总结micro服务如何运行在kubernetes 

## Overview

基础设施: consul apollo prometheus 

Go: go-micro gin 

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