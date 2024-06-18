package main

import (
	"os"
	"strconv"
)

func main() {

	redisAddr := os.Getenv("REDIS_ADDRESS")
	appPort, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		appPort = 8888
	}

	server := NewServer(
		ServerConfig{
			Port:      appPort,
			RedisAddr: redisAddr,
		},
	)

	server.Run()

}
