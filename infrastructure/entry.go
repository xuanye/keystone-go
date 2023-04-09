package infrastructure

import (
	"go.uber.org/dig"

	"github.com/xuanye/keystone-go/common"
	commonConf "github.com/xuanye/keystone-go/common/config"
	"github.com/xuanye/keystone-go/domain/contract"
	"github.com/xuanye/keystone-go/infrastructure/config"
)

func InitContainer(c *dig.Container) {

	c.Provide(config.NewViperConfigBuilder)
	c.Provide(func(builder contract.ConfigBuilder, environment *common.HostEnvironment) *commonConf.Settings {
		return builder.Build(environment)
	})
}
