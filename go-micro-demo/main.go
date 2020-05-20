//__author__ = "YaoYao"
//Date: 2020/5/4
package main

import (
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-plugins/registry/kubernetes/v2"
)

func main() {
	reg := kubernetes.NewRegistry()
	service := micro.NewService(
		micro.Registry(reg),
	)
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
