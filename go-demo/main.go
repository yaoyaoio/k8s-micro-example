//__author__ = "YaoYao"
//Date: 2020/4/29
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type Micro struct {
	Name    string `json:"name"`
	Version int    `json:"version"`
}

func main() {
	service := gin.Default()
	service.GET("/", func(ctx *gin.Context) {
		p := &Micro{Name: "micro", Version: 1.0}
		ctx.JSON(200, &p)
		return
	})
	service.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ok",
		})
		return
	})
	if err := service.Run(":8000"); err != nil {
		fmt.Println(err)
	}
}
