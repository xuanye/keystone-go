package domain

import (
	"github.com/xuanye/keystone-go/domain/service"
	"go.uber.org/dig"
)

func InitContainer(c *dig.Container) {
	c.Provide(service.NewApplicationService)
}
