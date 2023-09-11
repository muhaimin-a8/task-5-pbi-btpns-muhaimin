package app

import (
	"github.com/spf13/viper"
	"log"
)

func LoadConfig() {
	// load env
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
}
