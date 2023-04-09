package config

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/xuanye/keystone-go/common/config"
	"github.com/xuanye/keystone-go/domain/contract"
)

type ViperConfigBuilder struct {
}

func NewViperConfigBuilder() contract.ConfigBuilder {
	return &ViperConfigBuilder{}
}

func (builder *ViperConfigBuilder) Build(env string) config.Settings {

	// 加载配置文件
	if env == "" {
		viper.SetConfigName("config") // 配置文件名（无扩展名）
	} else {
		viper.SetConfigName(fmt.Sprintf("%s.config", env)) // 带有环境的
	}
	viper.SetConfigType("yaml") // 配置文件类型
	viper.AddConfigPath("./")   // 配置文件路径（当前目录）

	var settings config.Settings

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("read config error1: %v\n", err)
	}

	if err := viper.Unmarshal(&settings); err != nil {
		fmt.Printf("read config error2: %v\n", err)
	}
	return settings
}
