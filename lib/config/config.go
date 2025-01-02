package config

import (
	"log/slog"

	configdir "github.com/danielsoro/wordpress-cli/utils"
	"github.com/spf13/viper"
)

type Credentials struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type Config struct {
	Url         string      `mapstructure:"url"`
	Credentials Credentials `mapstructure:"credentials"`
}

var config Config

func NewConfig() Config {
	return NewConfigWithPath(configdir.LocalConfig("wordpress-cli"))
}

func NewConfigWithPath(path string) Config {
	viper.AddConfigPath(path)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		slog.Debug("Not able to find the config file")
	}

	if err := viper.Unmarshal(&config); err != nil {
		slog.Debug("Not possible to unmarshal config, verify if your configuration is correct")
	}

	return config
}
