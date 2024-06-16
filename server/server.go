package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/redis/go-redis/v9"
)

type Server struct {
	Address             string
	Port                int
	HostId              string
	Clients             map[string]*Client
	ClientHandler       *ClientHandler
	RedisMessageHandler *RedisMessageHandler
	RedisClient         *redis.Client
	mutex               sync.RWMutex
	ctx                 context.Context
}

type ServerConfig struct {
	Address     string
	Port        int
	HostID      string
	RedisClient *redis.Client
}

func NewServer(conf ServerConfig) *Server {
	ctx := context.Background()

	server := &Server{
		Address: conf.Address,
		Port:    conf.Port,
		HostId:  conf.HostID,
		Clients: make(map[string]*Client),
		ClientHandler: &ClientHandler{
			ComChan: make(chan Command),
		},
		RedisMessageHandler: &RedisMessageHandler{},
		RedisClient:         conf.RedisClient,
		mutex:               sync.RWMutex{},
		ctx:                 ctx,
	}
	server.ClientHandler.Server = server
	server.RedisMessageHandler.Server = server
	return server
}

func (server *Server) AddClient(client *Client) {
	server.mutex.Lock()
	defer server.mutex.Unlock()
	server.Clients[client.ID] = client
}

func (s *Server) Run() {
	fmt.Println("Initializing the chat server...")

	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Address, s.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer listener.Close()

	go s.ClientHandler.HandleClientMessage()
	go s.RedisMessageHandler.HandleRedisMessage()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to read connection from %s", conn.RemoteAddr().String())
			continue
		}
		log.Printf("New client from %s has join the server", conn.RemoteAddr().String())
		client := NewClient(conn, s.ClientHandler.ComChan)
		s.AddClient(client)
		go client.ReadInput()

	}

}
