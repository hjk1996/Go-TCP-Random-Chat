package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/redis/go-redis/v9"
)

func main() {

	fmt.Println("Initializing the chat server...")

	listener, err := net.Listen("tcp", "localhost:8888")
	defer listener.Close()

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	ctx := context.Background()
	redisClient := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379",
			Password: "1234",
			DB:       0,
		},
	)

	go func() {
		redisClient.Subscribe(ctx, "chat-message")
	}()

	server := NewServer()
	go server.Run()

	for {

		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to read connection from %s", conn.RemoteAddr().String())
			continue
		}

		log.Printf("New client %s has join the server", conn.RemoteAddr().String())

		client := NewClient(conn, server.ComChan)

		go client.readInput()

	}

}
