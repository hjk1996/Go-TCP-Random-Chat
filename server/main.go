package main

import "os"

func main() {

	redisAddr := os.Getenv("REDIS_ADDRESS")

	server := NewServer(
		ServerConfig{
			Port:      8888,
			RedisAddr: redisAddr,
		},
	)

	server.Run()

}
