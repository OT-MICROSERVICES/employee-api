package config

import (
	"employee-api/model"
	"github.com/spf13/viper"
)

// ReadConfigAndProperty is a method for getting the configuration file details
func ReadConfigAndProperty() model.Config {
	var config model.Config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/employee-api/")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return config
	}
	err = viper.Unmarshal(&config)
	return config
}
