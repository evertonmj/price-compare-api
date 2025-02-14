package configs_db

import (
    "github.com/redis/go-redis/v9"
)

func NewConnection() *redis.Client {
    client := redis.NewClient(&redis.Options{
        Addr:     "127.0.0.1:6379",
        Username: "default",
        Password: "secret",
    })
    return client
}