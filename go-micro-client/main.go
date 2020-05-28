//__author__ = "YaoYao"
//Date: 2020/5/28
package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	grpcc "github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/kubernetes/v2"
	proto "github.com/yaoliu/k8s-micro/proto"
	"time"
)

func main() {
	service := micro.NewService(
		micro.Name("go-micro-client"),
		micro.Client(makeMicroRPCClient()),
		micro.Registry(makeMicroRegistry()),
	)
	service.Init()
	registryRPCHandler(service.Client())
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func makeMicroRPCClient() client.Client {
	return grpcc.NewClient()
}

func makeMicroRegistry() registry.Registry {
	return kubernetes.NewRegistry()
}

func registryRPCHandler(s client.Client) {
	timer := time.NewTimer(time.Second * 10)
	for {
		select {
		case <-timer.C:
			greeter := proto.NewGreeterService("go.micro.server", s)
			rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "Yao"})
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(rsp.Greeting)
		}
	}
}
