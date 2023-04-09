package contract

import (
	"github.com/xuanye/keystone-go/common"
	"github.com/xuanye/keystone-go/common/config"
)

type ConfigBuilder interface {
	Build(env *common.HostEnvironment) *config.Settings
}
