//__author__ = "YaoYao"
//Date: 2020/5/4
package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	grpcc "github.com/micro/go-micro/v2/client/grpc"
	"github.com/micro/go-micro/v2/server"
	grpcs "github.com/micro/go-micro/v2/server/grpc"
	"github.com/micro/go-plugins/registry/kubernetes/v2"
	proto "github.com/yaoliu/k8s-micro/proto"
	_ "net/http/pprof"
)

type Greeter struct {
}

func (g Greeter) Hello(ctx context.Context, request *proto.HelloRequest, response *proto.HelloResponse) error {
	response.Greeting = "Hello " + request.Name + "!"
	return nil
}

var (
	DefaultServerPort = ":9100"
)

func main() {
	service := micro.NewService(
		micro.Name("go-micro-srv"),
		micro.Server(grpcs.NewServer(server.Address(DefaultServerPort), server.Name("go-micro-srv"))),
		micro.Client(grpcc.NewClient()),
		micro.Registry(kubernetes.NewRegistry()),
	)
	service.Init()

	_ = proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
