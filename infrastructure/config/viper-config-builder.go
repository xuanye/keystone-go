package config

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/xuanye/keystone-go/common"
	"github.com/xuanye/keystone-go/common/config"
	"github.com/xuanye/keystone-go/domain/contract"
)

type ViperConfigBuilder struct {
}

func NewViperConfigBuilder() contract.ConfigBuilder {
	return &ViperConfigBuilder{}
}

func (builder *ViperConfigBuilder) Build(env *common.HostEnvironment) *config.Settings {

	viper.SetDefault("port", 8080)

	viper.SetConfigName("config") // 配置文件名（无扩展名）
	viper.SetConfigType("yaml")   // 配置文件类型
	viper.AddConfigPath("./")     // 配置文件路径（当前目录）

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("error reading prod config file: %s", err))
	}

	if env.EnvironmentName != "" {
		envConfig := viper.New()
		envConfig.SetConfigName(fmt.Sprintf("config.%s", env.EnvironmentName))
		envConfig.SetConfigType("yaml")
		envConfig.AddConfigPath("./")
		err = envConfig.ReadInConfig()
		if err != nil {
			panic(fmt.Errorf("error reading %s config file: %s", env.EnvironmentName, err))
		}
		err = viper.MergeConfigMap(envConfig.AllSettings())
		if err != nil {
			panic(fmt.Errorf("error reading %s config file: %s", env.EnvironmentName, err))
		}
	}

	var settings config.Settings
	if err := viper.Unmarshal(&settings); err != nil {
		fmt.Printf("read config error2: %v\n", err)
	}

	return &settings
}
