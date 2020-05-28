//__author__ = "YaoYao"
//Date: 2020/5/28
package main

import (
	"fmt"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/kubernetes/v2"
	"net/http"
)

func main() {

	opts := []web.Option{
		web.Name("go-micro-web"),
		web.Registry(makeMicroRegistry()),
		web.Address(":9200"),
	}
	service := web.NewService(opts...)

	service.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("go-micro-web"))
	})

	if err := service.Init(); err != nil {
		fmt.Println(err)
	}

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func makeMicroRegistry() registry.Registry {
	return kubernetes.NewRegistry()
}
