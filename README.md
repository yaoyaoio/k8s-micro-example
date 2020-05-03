# k8s-micro

### 1. 一个简单的gin服务运行在k8s上的demo
    
    ```
    cd go-demo
    make build
    kubectl apply -f k8s-Deployment.yaml
    kubectl apply -f k8s-Service.yaml
    ```

### 2. 实现一个简单的nfs flexVolume 方便理解相关flexVolume插件使用

**介绍: Flexvolume提供了一种扩展k8s存储插件的方式，用户可以自定义自己的存储插件**
   
   ls flexvolume/yao~nfs/nfs.sh 官方提供的nfs实现方式
   ls flexvolume/yao~nfs/nfs.go 我写的一个go语言的demo

**Flexvolume接口**

   实现Flexvolume插件接口 包括 init/attach/detach/waitforattach/isattached/mountdevice/unmountdevice/mount/umount
   
   插件脚本需要放在node的节点里 k8s会自动watch对应/usr/libexec/kubernetes/kubelet-plugins/volume/exec/下的目录变化
   
    ```
    /usr/libexec/kubernetes/kubelet-plugins/volume/exec/yao~nfs/nfs mount <mount dir> <json param>
    ```
    
   pv创建方式如下
     
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
    
