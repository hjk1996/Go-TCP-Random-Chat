package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"example.com/chat/model"
)

type RedisMessageHandler struct {
	Server *Server
}

func (rh *RedisMessageHandler) HandleRedisMessage() {
	channel := fmt.Sprintf("channel:%s", rh.Server.HostId)
	sub := rh.Server.RedisClient.Subscribe(rh.Server.ctx, channel)
	defer sub.Close()
	ch := sub.Channel()
	for rawMessage := range ch {
		var msg model.BroadcastMessage
		json.Unmarshal([]byte(rawMessage.Payload), &msg)
		switch msg.MessageType {
		case model.BROADCAST_OPPONENT_JOIN_ROOM:
			go rh.handleOpponentJoinRoom(msg)
		case model.BROADCAST_OPPONENT_LEAVE_ROOM:
			go rh.handleLeaveRoom(msg)
		case model.BROADCAST_OPPONENT_SEND_MESSAGE:
			go rh.handleSendMessage(msg)
		}

	}
}

func (rh *RedisMessageHandler) handleOpponentJoinRoom(msg model.BroadcastMessage) {
	for _, targetID := range msg.Targets {
		rh.writeMessage(targetID, msg, model.CLIENT_OPPONENT_JOIN_ROOM)

	}

}

// 방에 남아있던 사람들에게 유저가 나갔다고 알려주고 방을 터트림
func (rh *RedisMessageHandler) handleLeaveRoom(msg model.BroadcastMessage) {
	rh.Server.mutex.Lock()
	defer rh.Server.mutex.Unlock()

	for _, targetID := range msg.Targets {
		target := rh.Server.Clients[targetID]
		if target == nil {
			continue
		}
		target.CurrentRoomId = ""
		rh.writeMessage(targetID, msg, model.CLIENT_OPPONENT_LEAVE_ROOM)

	}
}

func (rh *RedisMessageHandler) handleSendMessage(msg model.BroadcastMessage) {
	for _, targetID := range msg.Targets {
		rh.writeMessage(targetID, msg, model.CLIENT_OPPONENT_SEND_MESSAGE)
	}
}

// TODO: 0613
// 클라이언트에게 메시지 보내주는거 구현해야함
func (rh *RedisMessageHandler) writeMessage(clientId string, msg model.BroadcastMessage, messageType model.ClientMessageType) {
	target := rh.Server.Clients[clientId]
	if target == nil {
		log.Printf("User %s  does not exist on server", clientId)
		return
	}

	content := strings.TrimSpace(msg.Content)
	newMsg := &model.ClientMessage{
		MessageType: messageType,
		SenderID:    msg.SenderId,
		Content:     content,
		Timestamp:   time.Now(),
	}
	target.Conn.Write([]byte(newMsg.ToJson()))
}
