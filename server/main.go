package main

import (
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func main() {

	redisClient := redis.NewClient(
		&redis.Options{
			Addr: "localhost:6379",
			DB:   0,
		},
	)

	server := NewServer(
		ServerConfig{
			Address:     "localhost",
			Port:        8888,
			HostID:      uuid.New().String(),
			RedisClient: redisClient,
		},
	)

	server.Run()

}
