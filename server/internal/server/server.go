package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	Address             string
	Port                int
	HostId              string
	Users               map[string]*User
	UserHandler       *UserHandler
	RedisMessageHandler *RedisMessageHandler
	RedisClient         *redis.Client
	mutex               sync.RWMutex
	ctx                 context.Context
}

type ServerConfig struct {
	Port      int
	RedisAddr string
}

func NewServer(conf ServerConfig) *Server {
	ctx := context.Background()

	redisClient := redis.NewClient(
		&redis.Options{
			Addr: conf.RedisAddr,
			DB:   0,
		},
	)

	server := &Server{
		Port:   conf.Port,
		HostId: uuid.New().String(),
		Users:  make(map[string]*User),
		UserHandler: &UserHandler{
			ComChan: make(chan Command),
		},
		RedisMessageHandler: &RedisMessageHandler{},
		RedisClient:         redisClient,
		mutex:               sync.RWMutex{},
		ctx:                 ctx,
	}
	server.UserHandler.Server = server
	server.RedisMessageHandler.Server = server
	return server
}

func (server *Server) AddClient(client *User) {
	server.mutex.Lock()
	defer server.mutex.Unlock()
	server.Users[client.ID] = client
}

func (s *Server) Run() {
	fmt.Println("Initializing the chat server...")
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	defer listener.Close()

	go s.UserHandler.HandleClientMessage()
	go s.RedisMessageHandler.HandleRedisMessage()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to read connection from %s", conn.RemoteAddr().String())
			continue
		}
		log.Printf("New client from %s has join the server", conn.RemoteAddr().String())
		client := NewUser(conn, s.UserHandler.ComChan)
		s.AddClient(client)
		go client.ReadInput()
	}

}
