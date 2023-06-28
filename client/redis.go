package client

import (
	"employee-api/config"
	"github.com/redis/go-redis/v9"
)

// CreateRedisClient is a method for generating client of Redis
func CreateRedisClient() *redis.Client {
	config := config.ReadConfigAndProperty()
	return redis.NewClient(&redis.Options{
		Addr:     config.Redis.Host,
		Password: config.Redis.Password,
		DB:       config.Redis.Database,
	})
}
