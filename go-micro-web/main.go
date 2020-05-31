//__author__ = "YaoYao"
//Date: 2020/5/28
package main

import (
	"fmt"
	"github.com/micro/go-micro/v2/web"
	"github.com/micro/go-plugins/registry/kubernetes/v2"
	"net/http"
	"os"
)

func main() {

	opts := []web.Option{
		web.Name("go-micro-web"),
		web.Registry(kubernetes.NewRegistry()),
		web.Address(":9200"),
	}
	service := web.NewService(opts...)

	service.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		podName := os.Getenv("HOSTNAME")
		_, _ = writer.Write([]byte(podName))
	})

	if err := service.Init(); err != nil {
		fmt.Println(err)
	}

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
