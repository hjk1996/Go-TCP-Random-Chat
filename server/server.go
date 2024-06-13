package main

import (
	"context"
	"sync"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	HostId              string
	Clients             map[string]*Client
	ClientHandler       *ClientHandler
	RedisMessageHandler *RedisMessageHandler
	RedisClient         *redis.Client
	mutex               sync.RWMutex
	ctx                 context.Context
}

func NewServer(HostId string, redisClient *redis.Client) *Server {
	ctx := context.Background()

	server := &Server{
		HostId:  HostId,
		Clients: make(map[string]*Client),
		ClientHandler: &ClientHandler{
			ComChan: make(chan Command),
		},
		RedisMessageHandler: &RedisMessageHandler{},
		RedisClient:         redisClient,
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
	go s.ClientHandler.HandleClientMessage()
	go s.RedisMessageHandler.HandleRedisMessage()
}
