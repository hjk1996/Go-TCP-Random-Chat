package main

import (
	"encoding/json"
	"fmt"
	"log"
	"example.com/chat/model"
)

type RedisMessageHandler struct {
	Server *Server
}

func (rh *RedisMessageHandler) HandleRedisMessage() {

	sub := rh.Server.RedisClient.Subscribe(rh.Server.ctx, fmt.Sprintf("channel:%s", rh.Server.HostId))
	defer sub.Close()
	ch := sub.Channel()
	for rawMessage := range ch {
		var msg model.BroadcastMessage
		json.Unmarshal([]byte(rawMessage.Payload), &msg)
		switch msg.MessageType {
		case model.BROADCAST_LEAVE_ROOM:
			rh.handleLeaveRoom(msg)
		case model.BROADCAST_SEND_MESSAGE:
			rh.handleSendMessage(msg)
		}

	}
}

// 방에 남아있던 사람들에게 유저가 나갔다고 알려주고 방을 터트림
func (rh *RedisMessageHandler) handleLeaveRoom(msg model.BroadcastMessage) {
	rh.Server.mutex.Lock()
	defer rh.Server.mutex.Unlock()

	for _, targetID := range msg.Targets {
		target := rh.Server.Clients[targetID]
		if target == nil {
			log.Printf("User %s  does not exist on server", targetID)
			continue
		}
		target.CurrentRoomId = ""

		target.Conn.Write([]byte(msg.Content))
	}
}

func (rh *RedisMessageHandler) handleSendMessage(msg model.BroadcastMessage) {
	rh.Server.mutex.Lock()
	defer rh.Server.mutex.Unlock()
	for _, targetID := range msg.Targets {
		target := rh.Server.Clients[targetID]
		if target == nil {
			log.Printf("User %s  does not exist on server", targetID)
			continue
		}
		target.Conn.Write([]byte(fmt.Sprintf("<%s>: %s\n", msg.SenderId, msg.Content)))
	}
}

// TODO: 0613
// 클라이언트에게 메시지 보내주는거 구현해야함
func (rh *RedisMessageHandler) writeMessage(clientId string) {
}
