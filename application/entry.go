package application

import (
	"github.com/xuanye/keystone-go/application/feature/application"
	"go.uber.org/dig"
)

func InitContainer(c *dig.Container) {

	c.Provide(application.NewHandler)
}
