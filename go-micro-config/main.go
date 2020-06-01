//__author__ = "YaoYao"
//Date: 2020/5/31
package main

import (
	"fmt"
	"github.com/micro/go-micro/v2/config"
	"github.com/micro/go-micro/v2/config/source/env"
	"github.com/micro/go-plugins/config/source/configmap/v2"
)

var (
	DefaultNamespace = "go-micro"
)

func main() {
	if cfg, err := config.NewConfig(); err == nil {
		err = cfg.Load(
			env.NewSource(),
			configmap.NewSource(configmap.WithNamespace(DefaultNamespace)),
		)
		if err == nil {
			fmt.Println(cfg.Map())
		}
		fmt.Println(err)
	}
}
