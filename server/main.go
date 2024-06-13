package main

import (
	"fmt"
	"log"
	"net"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

func main() {

	fmt.Println("Initializing the chat server...")
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer listener.Close()

	redisClient := redis.NewClient(
		&redis.Options{
			Addr: "localhost:6379",
			DB:   0,
		},
	)

	hostId := uuid.New().String()
	server := NewServer(hostId, redisClient)

	server.Run()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to read connection from %s", conn.RemoteAddr().String())
			continue
		}
		log.Printf("New client from %s has join the server", conn.RemoteAddr().String())
		client := NewClient(conn, server.ClientHandler.ComChan)
		server.AddClient(client)
		go client.ReadInput()

	}

}
