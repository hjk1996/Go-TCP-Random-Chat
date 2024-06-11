package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	"example.com/chat/model"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type Server struct {
	HostId      string
	ComChan     chan Command
	Clients     map[string]*Client
	RedisClient *redis.Client
	mutex       sync.RWMutex
}

func NewServer(HostId string, redisClient *redis.Client) *Server {
	return &Server{
		HostId:      HostId,
		ComChan:     make(chan Command),
		Clients:     make(map[string]*Client),
		RedisClient: redisClient,
		mutex:       sync.RWMutex{},
	}
}

func (server *Server) AddClient(client *Client) {
	server.mutex.Lock()
	defer server.mutex.Unlock()
	server.Clients[client.ID] = client
}

func (server *Server) Run() {
	go server.handleClientMessage()
	go server.handleRedisMessage()
}

func (server *Server) handleClientMessage() {
	for cmd := range server.ComChan {
		switch cmd.CommandType {
		case CMD_JOIN_ROOM:
			server.joinRoom(cmd)
		case CMD_LEAVE_ROOM:
			server.leaveRoom(cmd)
		case CMD_NEW_ROOM:
			server.createRoom(cmd)
		case CMD_SEND_MESSAGE:
			server.sendMessage(cmd)
		}

	}
}

func (server *Server) handleRedisMessage() {
	ctx := context.Background()
	sub := server.RedisClient.Subscribe(ctx, fmt.Sprintf("channel:%s", server.HostId))
	defer sub.Close()
	ch := sub.Channel()
	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}
}

func (server *Server) joinRoom(cmd Command) {
	log.Println("join")
}

func (server *Server) leaveRoom(cmd Command) {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	if cmd.Client.CurrentRoomId == "" {
		log.Printf("Client %s does not belong to any room", cmd.Client.Conn.RemoteAddr().String())
		return
	}

	ctx := context.Background()
	var roomInfo model.RoomInfo
	infoString, err := server.RedisClient.Get(ctx, fmt.Sprintf("room:%s", cmd.Client.CurrentRoomId)).Result()
	if err != nil {
		log.Printf("Failed to get the information of the room %s", cmd.Client.CurrentRoomId)
		return
	}
	err = json.Unmarshal([]byte(infoString), &roomInfo)
	if err != nil {
		log.Printf("Failed to unmarshal room information. : %s", infoString)
		return
	}

	hostMap := make(map[string][]string)
	newClients := []model.ClientInfo{}
	for _, client := range roomInfo.Clients {
		if client.ID == cmd.Client.ID {
			continue
		}
		hostMap[client.HostID] = append(hostMap[client.HostID], client.ID)
		newClients = append(newClients, client)
	}

	newRoomInfo := model.RoomInfo{
		ID:      roomInfo.ID,
		Clients: newClients,
	}
	jsonData, err := json.Marshal(newRoomInfo)

	if err != nil {
		log.Printf("Failed to marshal json data")
		return
	}
	// 클라이언트 방 초기화
	cmd.Client.CurrentRoomId = ""
	// 방 정보 업데이트
	err = server.RedisClient.Set(ctx, fmt.Sprintf("room:%s", roomInfo.ID), jsonData, 0).Err()

	if err != nil {
		log.Printf("Failed to update the room %s information", roomInfo.ID)
		return
	}

	for hostId, clientIds := range hostMap {
		message := model.BroadcastMessage{
			MessageType: model.BROADCAST_LEAVE_ROOM,
			Targets:     clientIds,
			Content:     fmt.Sprintf("User %s has left the room.", cmd.Client.ID),
		}
		data, err := json.Marshal(message)

		if err != nil {
			log.Printf("Failed to send message to the host %s", hostId)
		}

		err = server.RedisClient.Publish(ctx, fmt.Sprintf("channel:%s", hostId), data).Err()

		if err != nil {
			log.Printf("Failed to send message to the host %s", hostId)
		}
	}

}

func (server *Server) createRoom(cmd Command) {
	server.mutex.Lock()
	defer server.mutex.Unlock()

	if cmd.Client.CurrentRoomId != "" {
		log.Printf("Client has already joined the room %s", cmd.Client.CurrentRoomId)
		return
	}

	// 고유햔 id 찾기
	ctx := context.Background()
	var roomId string
	for {
		roomId = uuid.New().String()
		_, err := server.RedisClient.Get(ctx, fmt.Sprintf("room:%s", roomId)).Result()
		if err == redis.Nil {
			break
		} else if err != nil {
			log.Printf("Failed to create the room: %s", err.Error())
			return
		} else {
			continue
		}

	}

	clients := make([]model.ClientInfo, 0, 2)
	clientInfo := model.ClientInfo{
		ID:     cmd.Client.ID,
		HostID: server.HostId,
	}
	clients = append(clients, clientInfo)
	roomInfo := model.RoomInfo{
		ID:      roomId,
		Clients: clients,
	}
	jsonData, err := json.Marshal(roomInfo)

	if err != nil {
		log.Printf("Failed to create the room: %s", err.Error())
		return
	}

	// 룸 정보 레디스에 삽입
	err = server.RedisClient.Set(ctx, fmt.Sprintf("room:%s", roomId), jsonData, 0).Err()
	if err != nil {
		log.Printf("Failed to set the room information: %s", err.Error())
		return
	}

	// 열려있는 방목록에 룸 정보 삽입
	err = server.RedisClient.RPush(ctx, "open_rooms", jsonData).Err()
	if err != nil {
		log.Printf("Failed to create the room: %s", err.Error())
		return
	}

	cmd.Client.CurrentRoomId = roomId

	log.Printf("Client opened the new room %s", roomId)

}

func (server *Server) sendMessage(cmd Command) {
	log.Println("msg")

}
