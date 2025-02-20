package configs_db

import (
	"github.com/redis/go-redis/v9"
)

// NewConnection creates and returns a new Redis client connection
func NewConnection() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379", // Redis server address
		Username: "default",        // Redis username
		Password: "secret",         // Redis password
	})
	return client
}
