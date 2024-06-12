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

func (s *Server) Run() {
	go s.handleClientMessage()
	go s.handleRedisMessage()
}

func (s *Server) handleClientMessage() {
	for cmd := range s.ComChan {
		switch cmd.CommandType {
		case CMD_JOIN_ROOM:
			s.joinRoom(cmd)
		case CMD_LEAVE_ROOM:
			s.leaveRoom(cmd)
		case CMD_NEW_ROOM:
			s.createRoom(cmd)
		case CMD_SEND_MESSAGE:
			s.sendMessage(cmd)
		case CMD_REMOVE_CLIENT:
			s.removeClient(cmd)
		case CMD_GET_CURRENT_ROOM:
			// TODO

		}
	}
}

func (s *Server) handleRedisMessage() {
	ctx := context.Background()
	sub := s.RedisClient.Subscribe(ctx, fmt.Sprintf("channel:%s", s.HostId))
	defer sub.Close()
	ch := sub.Channel()
	for rawMessage := range ch {
		var message model.BroadcastMessage
		json.Unmarshal([]byte(rawMessage.Payload), &message)
		switch message.MessageType {
		case model.BROADCAST_LEAVE_ROOM:
			for _, targetID := range message.Targets {
				target := s.Clients[targetID]
				if target == nil {
					log.Printf("User %s  does not exist on server", targetID)
					continue
				}
				target.CurrentRoomId = ""
				target.Conn.Write([]byte(message.Content))
			}
		}
	}
}

func (s *Server) joinRoom(cmd Command) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if cmd.Client.CurrentRoomId != "" {
		log.Printf("Client has already joined the room %s", cmd.Client.CurrentRoomId)
		return
	}

	ctx := context.Background()

	val, err := s.RedisClient.LPop(ctx, "open_rooms").Result()
	if err == redis.Nil {
		log.Println("No room exists. create new one")
		go s.createRoom(cmd)
		return
	}

	var roomInfo model.RoomInfo
	err = json.Unmarshal([]byte(val), &roomInfo)
	if err != nil {
		log.Printf("Failed to join a room: %s", err.Error())
		return
	}

	// 모종의 이유로 들어간 방에 나혼자 밖에 없을 때.
	// 그냥 방을 다시 만듬
	if len(roomInfo.Clients) == 0 {
		go s.createRoom(cmd)
		return
	}

	// 새로운 방정보 만들기
	newRoomInfo, err := roomInfo.Copy()
	if err != nil {
		log.Printf("Failed to join a room: %s", err.Error())
	}

	newRoomInfo.Clients = append(newRoomInfo.Clients, model.ClientInfo{
		ID:     cmd.Client.ID,
		HostID: s.HostId,
	})
	newRoomData, err := json.Marshal(newRoomInfo)
	if err != nil {
		log.Printf("Failed to join a room: %s", err.Error())
		return
	}

	c := s.Clients[cmd.Client.ID]
	if c == nil {
		log.Printf("Failed to find the client %s", cmd.Client.ID)
		return
	}
	c.CurrentRoomId = roomInfo.ID

	//레디스에 새로운 방정보 업데이트
	err = s.RedisClient.Set(ctx, fmt.Sprintf("room:%s", roomInfo.ID), newRoomData, 0).Err()

	if err != nil {
		log.Printf("Failed to join a room: %s", err.Error())
		return
	}

	// 기존에 방에 있던 클라이언트들에게 새로운 유저가 들어왔다고 알림
	for _, client := range roomInfo.Clients {
		joinMessage := fmt.Sprintf("User %s has joined the room", client.ID)
		if client.HostID == s.HostId {
			s.Clients[client.ID].Conn.Write([]byte(joinMessage))
		} else {
			go s.broadcastMessage(
				client.HostID,
				model.BROADCAST_JOIN_ROOM,
				[]string{client.ID},
				joinMessage)

		}
	}

}

func (s *Server) leaveRoom(cmd Command) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	// 나갈 방이 없을 때
	if cmd.Client.CurrentRoomId == "" {
		log.Printf("Client %s does not belong to any room", cmd.Client.Conn.RemoteAddr().String())
		return
	}

	ctx := context.Background()

	// 레디스에서 현재 유저 방의 정보를 받아옴.
	roomInfo, err := s.getRoomInfo(cmd.Client.CurrentRoomId)
	if err != nil {
		log.Printf("Failed to get room info: %s", err.Error())
		return
	}

	// 유저 혼자있는 방일 때 그냥 방을 삭제해버리고 함수 종료
	if len(roomInfo.Clients) < 2 {
		err := s.RedisClient.Del(ctx, fmt.Sprintf("room:%s", cmd.Client.CurrentRoomId)).Err()
		if err != nil {
			log.Printf("Failed to delete the room %s : %s", cmd.Client.CurrentRoomId, err.Error())
		}
		return
	}

	// 유저가 나갔다는 메시지 전달할 대상 탐색
	hostMap := make(map[string][]string)
	for _, client := range roomInfo.Clients {
		// 나간 당사자면 메세지 전달 목록에 포함 안함
		if client.ID == cmd.Client.ID {
			continue
		}
		hostMap[client.HostID] = append(hostMap[client.HostID], client.ID)
	}

	// 클라이언트 방 초기화
	cmd.Client.CurrentRoomId = ""
	// 방 삭제
	err = s.RedisClient.Del(ctx, fmt.Sprintf("room:%s", cmd.Client.CurrentRoomId)).Err()
	if err != nil {
		log.Printf("Failed to delete the room %s : %s", cmd.Client.CurrentRoomId, err.Error())
	}

	// 메세지 브로드캐스팅
	// TODO: 이후에 방에 있던 클라이언트의 방 ID도 초기해줘야함.
	for hostId, clientIds := range hostMap {
		go s.broadcastMessage(hostId, model.BROADCAST_LEAVE_ROOM, clientIds, fmt.Sprintf("User %s has left the room.", cmd.Client.ID))
	}

}

func (s *Server) createRoom(cmd Command) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if cmd.Client.CurrentRoomId != "" {
		log.Printf("Client has already joined the room %s", cmd.Client.CurrentRoomId)
		return
	}

	// 고유햔 id 찾기
	ctx := context.Background()
	var roomId string
	for {
		roomId = uuid.New().String()
		// 방정보 조회
		_, err := s.getRoomInfo(roomId)
		if err == redis.Nil {
			break
		} else if err != nil {
			log.Printf("Failed to create the room: %s", err.Error())
			return
		} else {
			continue
		}

	}

	clients := make([]model.ClientInfo, 0, 1)
	clientInfo := model.ClientInfo{
		ID:     cmd.Client.ID,
		HostID: s.HostId,
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
	err = s.RedisClient.Set(ctx, fmt.Sprintf("room:%s", roomId), jsonData, 0).Err()
	if err != nil {
		log.Printf("Failed to set the room information: %s", err.Error())
		return
	}

	// 열려있는 방목록에 룸 정보 삽입
	err = s.RedisClient.RPush(ctx, "open_rooms", jsonData).Err()
	if err != nil {
		log.Printf("Failed to create the room: %s", err.Error())
		return
	}

	cmd.Client.CurrentRoomId = roomId

	log.Printf("Client opened the new room %s", roomId)

}

func (s *Server) sendMessage(cmd Command) {
	log.Println("msg")

}

func (s *Server) removeClient(cmd Command) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	ctx := context.Background()
	if cmd.Client.CurrentRoomId == "" {
		delete(s.Clients, cmd.Client.ID)
		log.Printf("Removed client %s", cmd.Client.ID)
		return
	}

	err := s.RedisClient.Del(ctx, fmt.Sprintf("room:%s", cmd.Client.CurrentRoomId)).Err()
	if err != nil {
		log.Printf("Failed to delete the room that the client left. :%s", err.Error())
	}
	delete(s.Clients, cmd.Client.ID)

}

func (s *Server) getRoomInfo(roomId string) (*model.RoomInfo, error) {
	ctx := context.Background()
	var roomInfo model.RoomInfo
	infoString, err := s.RedisClient.Get(ctx, fmt.Sprintf("room:%s", roomId)).Result()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(infoString), &roomInfo)
	if err != nil {
		return nil, err
	}

	return &roomInfo, nil

}

func (s *Server) broadcastMessage(
	channel string, messageType int, targets []string, content string) {
	ctx := context.Background()

	message := model.BroadcastMessage{
		MessageType: messageType,
		Targets:     targets,
		Content:     content,
	}
	data, err := json.Marshal(message)

	if err != nil {
		log.Printf("Failed to send message to channel %s", channel)
		return
	}

	err = s.RedisClient.Publish(ctx, fmt.Sprintf("channel:%s", channel), data).Err()

	if err != nil {
		log.Printf("Failed to send message to channel %s", channel)
		return
	}

}
