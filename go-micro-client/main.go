//__author__ = "YaoYao"
//Date: 2020/5/28
package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	grpcc "github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-plugins/registry/kubernetes/v2"
	proto "github.com/yaoliu/k8s-micro/proto"
	"time"
)

var (
	DefaultServiceName = "go-micro-client"
	DefaultSrvName     = "go-micro-srv"
)

func main() {
	service := micro.NewService(
		micro.Name(DefaultServiceName),
		micro.Client(grpcc.NewClient()),
		micro.Registry(kubernetes.NewRegistry()),
	)
	service.Init()
	go func() {
		registryRPCHandler(service.Client())
	}()
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}

}

func registryRPCHandler(s client.Client) {
	timer := time.NewTicker(time.Second * 10)
	defer timer.Stop()
	for {
		select {
		case <-timer.C:
			greeter := proto.NewGreeterService(DefaultSrvName, s)
			rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "Yao"})
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Server", rsp.Greeting)
			}
		}
	}
}
