package client

import (
	"employee-api/config"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

// CreateRedisClient is a method for generating client of Redis
func CreateRedisClient() *redis.Client {
	config, err := config.ReadConfigAndProperty()
	if err != nil {
		logrus.Errorf("Unable to read the configuration file: %v", err)
	}
	return redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host,
		Password: config.Redis.Password,
		DB:       config.Redis.Database,
	})
}
