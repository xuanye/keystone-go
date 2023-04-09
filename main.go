package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuanye/keystone-go/adapter"
	"github.com/xuanye/keystone-go/application"
	"github.com/xuanye/keystone-go/common"
	"github.com/xuanye/keystone-go/common/config"
	"github.com/xuanye/keystone-go/domain"
	"github.com/xuanye/keystone-go/infrastructure"
	"go.uber.org/dig"
	"os"
)

func setupEnvironment() *common.HostEnvironment {
	env := os.Getenv("GO_ENV")
	if env == "" {
		env = os.Getenv("APP_ENV")
	}

	return &common.HostEnvironment{EnvironmentName: env}
}

func setDependency() *dig.Container {

	container := dig.New()

	container.Provide(setupEnvironment)

	adapter.InitContainer(container)
	application.InitContainer(container)
	domain.InitContainer(container)
	infrastructure.InitContainer(container)

	return container
}

func main() {

	c := setDependency()

	// 运行容器

	if err := c.Invoke(func(router *gin.Engine, settings *config.Settings) {
		err := router.Run(fmt.Sprintf(":%d", settings.Port))
		if err != nil {
			panic(err)
		}
	}); err != nil {
		panic(err)
	}

}
