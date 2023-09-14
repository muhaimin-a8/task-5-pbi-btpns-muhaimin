package app

import (
	"github.com/spf13/viper"
)

func LoadConfig(configPath string) error {
	// load env
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
