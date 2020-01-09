package providers

import (
	"L/config"
	"github.com/jinzhu/configor"
)

type Config struct {
	config.App  // 应用配置
	config.Http // Http 服务配置
}

func NewConfig() *Config {
	Config := Config{}
	configor.Load(&Config, "config.yml") // 加载配置文件
	return &Config
}
