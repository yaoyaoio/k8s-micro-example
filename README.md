# k8s-micro

用于总结micro服务如何运行在kubernetes 包括go-micro gin consul apollo prometheus等组件及基础设施

## Overview


## Features 


## Getting Started

- [Gin on Kubernetes Demo]()
- [Simple NFS flexVolume]()
- [Go-micro on Kubernetes Demo]()

### Gin on Kubernetes Demo
    
```
cd go-gin-demo
make build
kubectl apply -f k8s-Deployment.yaml
kubectl apply -f k8s-Service.yaml
```

### Simple NFS flexVolume

**1) 介绍: Flexvolume提供了一种扩展k8s存储插件的方式，用户可以自定义自己的存储插件**

```   
ls flexvolume/yao~nfs/nfs.sh 官方提供的nfs实现方式
ls flexvolume/yao~nfs/nfs.go 我写的一个go语言的demo
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
    
### Go-micro on Kubernetes Demo