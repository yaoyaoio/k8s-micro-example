# k8s-micro

用于总结micro服务如何运行在kubernetes 

## Overview

基础设施: consul apollo prometheus 

Go: go-micro gin 

Java: 待定

Python: 待定

不断连载 包你满意 

## Features 


## Getting Started

- [Gin on Kubernetes Demo]()
- [Simple NFS flexVolume]()
- [Go-micro(Web) on Kubernetes Demo]()
- [Go-micro(RPC) on Kubernetes Demo]()
- [Promethues on Kubernetes Demo]()
- [Apollo on Kubernetes Demo]()
- [Consul on Kubernetes Demo]()
- [Simple NS Docker Demo]()

### Gin on Kubernetes Demo
    
```
cd go-gin-demo
make build or docker pull liuyao/go-gin-demo
kubectl apply -f k8s-PersistentVolumeClaim.yaml 
kubectl apply -f k8s-Deployment.yaml
kubectl apply -f k8s-Service.yaml
```

### Simple NFS flexVolume

**1) 介绍: Flexvolume提供了一种扩展k8s存储插件的方式，用户可以自定义自己的存储插件**

```   
ls flexvolume/yao~nfs/nfs.sh 官方提供的nfs shell实现方式
ls flexvolume/yao~nfs/nfs.go 我写的 go 实现方式
```

**2) Flexvolume接口**

实现Flexvolume插件接口 包括 
   `init/attach/detach/waitforattach/isattached/mountdevice/unmountdevice/mount/umount`
   
插件脚本需要放在node的节点里 k8s会自动watch对应/usr/libexec/kubernetes/kubelet-plugins/volume/exec/下的目录变化
   
```
/usr/libexec/kubernetes/kubelet-plugins/volume/exec/yao~nfs/nfs mount <mount dir> <json param>
```
    
**3)pv创建方式如下**
     
```  
apiVersion: v1
kind: PersistentVolume
metadata:
  name: pv-flex-nfs
spec:
  capacity:
    storage: 10Gi
  accessModes:
    - ReadWriteMany
  flexVolume:
    driver: "yao/nfs"
    fsType: "nfs"
    options:
      server: "192.168.0.114" # 改成你自己的NFS服务器地址
      share: "export"
```      
 
###  Simple NS Docker Demo

一个简单的使用Go syscall 来调用namespace 进入容器状态的demo
    
### Go-micro on Kubernetes Demo