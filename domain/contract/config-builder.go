package contract

import "github.com/xuanye/keystone-go/common/config"

type ConfigBuilder interface {
	Build(env string) config.Settings
}
