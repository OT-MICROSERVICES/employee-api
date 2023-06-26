package config

import (
	"employee-api/model"
	"github.com/spf13/viper"
)

// ReadConfigAndProperty is a method for getting the config file details
func ReadConfigAndProperty() (model.Config, error) {
	var config model.Config
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("/etc/employee-api/")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		return config, err
	}
	err = viper.Unmarshal(&config)
	return config, nil
}
