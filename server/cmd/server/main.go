package main

import (
	"os"
	"strconv"
	"example.com/chat/internal/server"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	redisAddr := os.Getenv("REDIS_ADDRESS")
	appPort, err := strconv.Atoi(os.Getenv("APP_PORT"))
	if err != nil {
		appPort = 8888
	}

	server := server.NewServer(
		server.ServerConfig{
			Port:      appPort,
			RedisAddr: redisAddr,
		},
	)

	server.Run()

}
