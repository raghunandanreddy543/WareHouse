package config

import (
	"github.com/spf13/viper"
)

func New(fileName string) *AppConfig {
	viper.SetConfigFile(fileName)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			// Config file was found but another error was produced
		}
	}

	config := LoadConfig()
	return config
}