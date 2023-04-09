package main

import (
	"github.com/gin-gonic/gin"
	"github.com/xuanye/keystone-go/adapter"
	"github.com/xuanye/keystone-go/application"
	"github.com/xuanye/keystone-go/domain"
	"go.uber.org/dig"
)

func setDependency() *dig.Container {

	container := dig.New()

	adapter.InitContainer(container)
	application.InitContainer(container)
	domain.InitContainer(container)
	return container
}

func main() {

	c := setDependency()

	// 运行容器

	if err := c.Invoke(func(router *gin.Engine) {
		err := router.Run(":8080")
		if err != nil {
			panic(err)
		}
	}); err != nil {
		panic(err)
	}

}
