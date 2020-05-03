//__author__ = "YaoYao"
//Date: 2020/5/2
package main

import (
	"fmt"
	"os"
)

type Driver interface {
	Init() (map[string]interface{}, error)                               //初始化存储插件时调用  插件需要返回是否需要要 attach 和 detach 操作
	Attach(jsonOptions, nodeName string) (map[string]interface{}, error) //
	Detach(mountDev, nodeName string) (map[string]interface{}, error)
	WaitForAttach(mountDev, jsonOptions string) (map[string]interface{}, error)
	IsAttached(jsonOptions, nodeName string) (map[string]interface{}, error)
	Mount(targetMountDir, jsonOptions string) (map[string]interface{}, error) //将存储卷挂载到指定目录中
	// /usr/libexec/kubernetes/kubelet-plugins/volume/exec/yao~nfs/nfs mount <mount dir> <json options>
	Unmount(targetMountDir string) (map[string]interface{}, error) //将存储卷取消挂载
}

type NFSFlexVolumeDriver struct{}

func newNFS() Driver {
	return &NFSFlexVolumeDriver{}
}

func (nfs NFSFlexVolumeDriver) Init() (map[string]interface{}, error) {
	panic("implement me")
}

func (nfs NFSFlexVolumeDriver) Attach(jsonOptions, nodeName string) (map[string]interface{}, error) {
	panic("implement me")
}

func (nfs NFSFlexVolumeDriver) Detach(mountDev, nodeName string) (map[string]interface{}, error) {
	panic("implement me")
}

func (nfs NFSFlexVolumeDriver) WaitForAttach(mountDev, jsonOptions string) (map[string]interface{}, error) {
	panic("implement me")
}

func (nfs NFSFlexVolumeDriver) IsAttached(jsonOptions, nodeName string) (map[string]interface{}, error) {
	panic("implement me")
}

func (nfs NFSFlexVolumeDriver) Mount(targetMountDir, jsonOptions string) (map[string]interface{}, error) {
	panic("implement me")
}

func (nfs NFSFlexVolumeDriver) Unmount(targetMountDir string) (map[string]interface{}, error) {
	panic("implement me")
}

type FlexVolume struct {
	Driver
}

func newFlexVolume(d Driver) FlexVolume {
	return FlexVolume{d}
}

func (fv *FlexVolume) Run(args []string) {
	if len(args) == 0 {

	}
	op := args[1]
	switch op {
	case "init":
		fmt.Println(1)
	}
}

func main() {
	fv := newFlexVolume(newNFS())
	fv.Run(os.Args)

}
