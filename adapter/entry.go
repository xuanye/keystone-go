package adapter

import (
	"github.com/xuanye/keystone-go/adapter/api"
	"go.uber.org/dig"
)

func InitContainer(c *dig.Container) {

	c.Provide(api.NewApplicationController, dig.Group("controller"))
	c.Provide(api.NewOrganizationController, dig.Group("controller"))
	c.Provide(NewRouterV1)
}
