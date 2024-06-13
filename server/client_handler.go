package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"example.com/chat/model"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

type ClientHandler struct {
	Server  *Server
	ComChan chan Command
}

func (ch *ClientHandler) HandleClientMessage() {
	for cmd := range ch.ComChan {
		switch cmd.CommandType {
		case CMD_JOIN_ROOM:
			ch.joinRoom(cmd)
		case CMD_LEAVE_ROOM:
			ch.leaveRoom(cmd)
		case CMD_NEW_ROOM:
			ch.createRoom(cmd)
		case CMD_SEND_MESSAGE:
			ch.sendMessageToOthers(cmd)
		case CMD_REMOVE_CLIENT:
			ch.removeClient(cmd)
		case CMD_GET_CURRENT_ROOM:

		}
	}
}

func (ch *ClientHandler) joinRoom(cmd Command) {
	ch.Server.mutex.Lock()
	defer ch.Server.mutex.Unlock()

	if cmd.Client.CurrentRoomId != "" {
		content := fmt.Sprintf("Client has already joined the room %s", cmd.Client.CurrentRoomId)
		ch.writeLogAndSendError(cmd.Client.ID, content)
		return
	}

	val, err := ch.Server.RedisClient.LPop(ch.Server.ctx, "open_rooms").Result()
	if err == redis.Nil {
		log.Println("No room exists. create new one")
		go ch.createRoom(cmd)
		return
	}

	var roomInfo model.RoomInfo
	err = json.Unmarshal([]byte(val), &roomInfo)
	if err != nil {
		content := fmt.Sprintf("Failed to join a room : %s", err.Error())
		ch.writeLogAndSendError(cmd.Client.ID, content)
		return
	}

	// 모종의 이유로 들어간 방에 나혼자 밖에 없을 때.
	// 그냥 방을 다시 만듬
	if len(roomInfo.Clients) == 0 {
		go ch.createRoom(cmd)
		return
	}

	// 새로운 방정보 만들기
	newRoomInfo, err := roomInfo.Copy()
	if err != nil {
		content := fmt.Sprintf("Failed to join a room : %s", err.Error())
		ch.writeLogAndSendError(cmd.Client.ID, content)
		return
	}

	newRoomInfo.Clients = append(newRoomInfo.Clients, model.ClientInfo{
		ID:     cmd.Client.ID,
		HostID: ch.Server.HostId,
	})

	c := ch.Server.Clients[cmd.Client.ID]
	if c == nil {
		log.Printf("Failed to find the client %s\n", cmd.Client.ID)
		return
	}
	c.CurrentRoomId = roomInfo.ID

	//레디스에 새로운 방정보 업데이트
	err = ch.Server.RedisClient.Set(
		ch.Server.ctx,
		fmt.Sprintf("room:%s",
			roomInfo.ID),
		newRoomInfo.ToJson(),
		0).Err()

	if err != nil {
		content := fmt.Sprintf("Failed to join a room : %s", err.Error())
		ch.writeLogAndSendError(cmd.Client.ID, content)
		return
	}

	log.Printf("User %s has joined the room %s\n", cmd.Client.ID, roomInfo.ID)
	// 기존에 방에 있던 클라이언트들에게 새로운 유저가 들어왔다고 알림
	for _, client := range roomInfo.Clients {
		go ch.broadcastMessage(
			client.HostID,
			model.BROADCAST_JOIN_ROOM,
			cmd.Client.ID,
			[]string{client.ID},
			"",
		)
	}

}

func (ch *ClientHandler) leaveRoom(cmd Command) {
	ch.Server.mutex.Lock()
	defer ch.Server.mutex.Unlock()

	// 나갈 방이 없을 때
	if cmd.Client.CurrentRoomId == "" {
		content := fmt.Sprintf("Client %s does not belng to any room", cmd.Client.ID)
		ch.writeLogAndSendError(cmd.Client.ID, content)
		return
	}

	// 레디스에서 현재 유저 방의 정보를 받아옴.
	roomInfo, err := ch.getRoomInfo(cmd.Client.CurrentRoomId)
	if err != nil {
		content := fmt.Sprintf("Failed to get room info : %s", err.Error())
		ch.writeLogAndSendError(cmd.Client.ID, content)
		return
	}

	// 유저 혼자있는 방일 때 그냥 방을 삭제해버리고 함수 종료
	if len(roomInfo.Clients) < 2 {
		err := ch.Server.RedisClient.Del(ch.Server.ctx, fmt.Sprintf("room:%s", cmd.Client.CurrentRoomId)).Err()
		if err != nil {
			content := fmt.Sprintf("Failed to delete the room %s : %s\n", cmd.Client.CurrentRoomId, err.Error())
			ch.writeLogAndSendError(cmd.Client.ID, content)
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
	err = ch.Server.RedisClient.Del(ch.Server.ctx, fmt.Sprintf("room:%s", cmd.Client.CurrentRoomId)).Err()
	if err != nil {
		content := fmt.Sprintf("Failed to delete the room : %s", err.Error())
		ch.writeLogAndSendError(cmd.Client.ID, content)
	}

	log.Printf("User %s has left the room %s\n", cmd.Client.ID, roomInfo.ID)

	// TODO: 이후에 방에 있던 클라이언트의 방 ID도 초기해줘야함.
	for hostId, clientIds := range hostMap {
		// 메세지 브로드캐스팅
		go ch.broadcastMessage(hostId,
			model.BROADCAST_LEAVE_ROOM,
			cmd.Client.ID,
			clientIds,
			fmt.Sprintf("User %s has left the room.\n", cmd.Client.ID),
		)
	}

}

func (ch *ClientHandler) createRoom(cmd Command) {
	ch.Server.mutex.Lock()
	defer ch.Server.mutex.Unlock()

	if cmd.Client.CurrentRoomId != "" {
		content := fmt.Sprintf("Client has already joined the room %s", cmd.Client.CurrentRoomId)
		ch.writeLogAndSendError(cmd.Client.ID, content)
		return
	}

	// 고유햔 id 찾기
	var roomId string
	for {
		roomId = uuid.New().String()
		// 방정보 조회
		_, err := ch.getRoomInfo(roomId)
		if err == redis.Nil {
			break
		} else if err != nil {
			content := fmt.Sprintf("Failed to create the room : %s", err.Error())
			ch.writeLogAndSendError(cmd.Client.ID, content)
			return
		} else {
			continue
		}

	}

	clients := make([]model.ClientInfo, 0, 1)
	clientInfo := model.ClientInfo{
		ID:     cmd.Client.ID,
		HostID: ch.Server.HostId,
	}
	clients = append(clients, clientInfo)
	roomInfo := model.RoomInfo{
		ID:      roomId,
		Clients: clients,
	}
	jsonData, err := json.Marshal(roomInfo)

	if err != nil {
		content := fmt.Sprintf("Failed to create the room: %s", err.Error())
		ch.writeLogAndSendError(cmd.Client.ID, content)
		return
	}

	// 룸 정보 레디스에 삽입
	err = ch.Server.RedisClient.Set(ch.Server.ctx, fmt.Sprintf("room:%s", roomId), jsonData, 0).Err()
	if err != nil {
		content := fmt.Sprintf("Failed to set the room information: %s", err.Error())
		ch.writeLogAndSendError(cmd.Client.ID, content)
		return
	}

	// 열려있는 방목록에 룸 정보 삽입
	err = ch.Server.RedisClient.RPush(ch.Server.ctx, "open_rooms", jsonData).Err()
	if err != nil {
		content := fmt.Sprintf("Failed to create the room: %s", err.Error())
		ch.writeLogAndSendError(cmd.Client.ID, content)
		return
	}

	cmd.Client.CurrentRoomId = roomId

	log.Printf("Client opened the new room %s\n", roomId)

}

func (ch *ClientHandler) sendMessageToOthers(cmd Command) {
	ch.Server.mutex.Lock()
	defer ch.Server.mutex.Unlock()

	if cmd.Client.CurrentRoomId == "" {
		content := fmt.Sprintf("User %s does not belong to any room.", cmd.Client.ID)
		ch.writeLogAndSendError(cmd.Client.ID, content)
		return
	}

	var roomInfo model.RoomInfo

	data, err := ch.Server.RedisClient.Get(ch.Server.ctx, fmt.Sprintf("room:%s", cmd.Client.CurrentRoomId)).Result()

	if err != nil {
		content := fmt.Sprintf("Failed to get the data from the redis. : %s", err.Error())
		ch.writeLogAndSendError(cmd.Client.ID, content)
		return
	}

	err = json.Unmarshal([]byte(data), &roomInfo)

	if err != nil {
		content := fmt.Sprintf("Failed to parse the json string. : %s", err.Error())
		ch.writeLogAndSendError(cmd.Client.ID, content)
	}

	if len(roomInfo.Clients) == 1 {
		return
	}

	for _, c := range roomInfo.Clients {
		if c.ID == cmd.Client.ID {
			continue
		}

		msg := strings.Join(cmd.Args, " ") + "\n"
		ch.broadcastMessage(c.HostID,
			model.BROADCAST_SEND_MESSAGE,
			cmd.Client.ID,
			[]string{c.ID},
			msg,
		)

	}

}
func (ch *ClientHandler) broadcastMessage(
	channel string, messageType model.BroadcastMessageType, senderId string, targets []string, content string) {

	message := &model.BroadcastMessage{
		MessageType: messageType,
		SenderId:    senderId,
		Targets:     targets,
		Content:     content,
	}

	err := ch.Server.RedisClient.Publish(ch.Server.ctx, fmt.Sprintf("channel:%s", channel), message.ToJson()).Err()

	if err != nil {
		log.Printf("Failed to send message to channel %s\n", channel)
		return
	}

}

func (ch *ClientHandler) removeClient(cmd Command) {
	ch.Server.mutex.Lock()
	defer ch.Server.mutex.Unlock()

	if cmd.Client.CurrentRoomId == "" {
		delete(ch.Server.Clients, cmd.Client.ID)
		log.Printf("Removed client %s", cmd.Client.ID)
		return
	}

	err := ch.Server.RedisClient.Del(ch.Server.ctx, fmt.Sprintf("room:%s", cmd.Client.CurrentRoomId)).Err()
	if err != nil {
		log.Printf("Failed to delete the room that the client left. :%s\n", err.Error())
	}
	delete(ch.Server.Clients, cmd.Client.ID)

}

func (ch *ClientHandler) getRoomInfo(roomId string) (*model.RoomInfo, error) {
	var roomInfo model.RoomInfo
	infoString, err := ch.Server.RedisClient.Get(ch.Server.ctx, fmt.Sprintf("room:%s", roomId)).Result()
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(infoString), &roomInfo)
	if err != nil {
		return nil, err
	}

	return &roomInfo, nil

}

func (ch *ClientHandler) sendMessageToClient(target string, messageType model.ClientMessageType, senderId string, content string) {
	msg := &model.ClientMessage{
		MessageType: messageType,
		SenderID:    senderId,
		Content:     content,
		Timestamp:   time.Now(),
	}
	client := ch.Server.Clients[target]

	if client == nil {
		log.Printf("Failed to find the client %s", target)
	} else {
		client.Conn.Write(msg.ToJson())
	}
}

func (ch *ClientHandler) writeLogAndSendError(target string, content string) {
	log.Printf(content + "\n")
	ch.sendMessageToClient(target, model.CLIENT_ERROR, target, content)

}
