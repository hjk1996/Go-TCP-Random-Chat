package server

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"example.com/chat/data"
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
			go ch.handleJoinRoom(cmd)
		case CMD_LEAVE_ROOM:
			go ch.handleLeaveRoom(cmd)
		case CMD_NEW_ROOM:
			go ch.handleCreateRoom(cmd)
		case CMD_SEND_MESSAGE:
			go ch.handleSendMessageToOpponent(cmd)
		case CMD_REMOVE_CLIENT:
			go ch.handleRemoveClient(cmd)
		case CMD_QUIT:
			ch.handleQuit(cmd)
		}
	}
}

func (ch *ClientHandler) handleJoinRoom(cmd Command) {
	ch.Server.mutex.Lock()
	defer ch.Server.mutex.Unlock()

	// 클라이언트가 이미 방에 들어가있으면 오류 발생시킴
	if cmd.Client.CurrentRoomId != "" {
		content := fmt.Sprintf("Client has already joined the room %s", cmd.Client.CurrentRoomId)
		ch.writeLogAndSendError(cmd.Client.ID, content)
		return
	}

	val, err := ch.Server.RedisClient.LPop(ch.Server.ctx, "open_rooms").Result()
	// join 명령어를 했늗네 들어갈 수 있는 방이 없으면 새로운 방을 만드는 goroutine을 생성
	if err == redis.Nil {
		log.Println("No room exists. create new one")
		go ch.handleCreateRoom(cmd)
		return
	}

	var roomInfo model.RoomInfo
	err = json.Unmarshal([]byte(val), &roomInfo)
	if err != nil {
		content := fmt.Sprintf("Failed to join a room : %s", err.Error())
		ch.writeLogAndSendError(cmd.Client.ID, content)
		return
	}

	// 모종의 이유로 들어간 방에 아무도 없을 때.
	// 그냥 방을 다시 만듬
	if len(roomInfo.Clients) == 0 {
		log.Println("No one is in the room. Create a new room")
		go ch.handleCreateRoom(cmd)
		return
	}

	// 내가 만든 방에 내가 들어갔을 때 방을 다시 만듬
	for _, client := range roomInfo.Clients {
		if cmd.Client.ID == client.ID {
			log.Println("User is trying to join the room that the user created")
			go ch.handleCreateRoom(cmd)
			return
		}
	}

	// 새로운 방정보 만들기고 레디스에 업데이트
	newRoomInfo := roomInfo.Copy()
	newRoomInfo.Clients = append(newRoomInfo.Clients, model.ClientInfo{
		ID:     cmd.Client.ID,
		HostID: ch.Server.HostId,
	})

	if err := ch.setRoomInRedis(newRoomInfo); err != nil {
		ch.writeLogAndSendError(cmd.Client.ID, err.Error())
		return

	}

	c := ch.Server.Clients[cmd.Client.ID]
	if c == nil {
		log.Printf("Failed to find the client %s\n", cmd.Client.ID)
		return
	}

	// 메모리에서 클라이언트 현재 방 ID 업데이트
	c.CurrentRoomId = roomInfo.ID

	ch.sendMessageToClient(
		cmd.Client.ID,
		model.CLIENT_JOIN_ROOM_CONFIRM,
		cmd.Client.ID,
		"",
	)

	log.Printf("User %s has joined the room %s\n", cmd.Client.ID, roomInfo.ID)
	// 기존에 방에 있던 클라이언트들에게 새로운 유저가 들어왔다고 알림
	for _, client := range roomInfo.Clients {
		go ch.broadcastMessage(
			client.HostID,
			model.BROADCAST_OPPONENT_JOIN_ROOM,
			cmd.Client.ID,
			[]string{client.ID},
			"",
		)
	}

}

func (ch *ClientHandler) handleLeaveRoom(cmd Command) {
	ch.Server.mutex.Lock()
	defer ch.Server.mutex.Unlock()

	// 클라이언트가 아무 방에 속해있지 않을 때
	if !ch.checkClientHasRoom(cmd) {
		ch.writeLogAndSendError(cmd.Client.ID, "Client does not belong to any room")
		return
	}

	// 레디스에서 현재 유저 방의 정보를 받아옴.
	roomInfo, err := ch.getRoomInfo(cmd.Client.CurrentRoomId)
	if err != nil {
		cmd.Client.CurrentRoomId = ""
		ch.writeLogAndSendError(cmd.Client.ID, fmt.Sprintf("Failed to get room info : %s", err.Error()))
		return
	}

	// 클라이언트 방 초기화
	cmd.Client.CurrentRoomId = ""

	// 유저 혼자있는 방일 때 그냥 방을 삭제해버리고 함수 종료
	if len(roomInfo.Clients) < 2 {
		ch.sendMessageToClient(
			cmd.Client.ID,
			model.CLIENT_LEAVE_ROOM_CONFIRM,
			cmd.Client.ID,
			"",
		)
		if err := ch.deleteRoomFromRedis(roomInfo.ID); err != nil {
			ch.writeLogAndSendError(cmd.Client.ID, fmt.Sprintf("Failed to delete the room %s : %s", cmd.Client.CurrentRoomId, err.Error()))
		}
		log.Printf("User %s has left the room %s\n", cmd.Client.ID, roomInfo.ID)

		return
	}

	// 클라이언트에게 방을 나갔다고 메시지 보냄
	ch.sendMessageToClient(
		cmd.Client.ID,
		model.CLIENT_LEAVE_ROOM_CONFIRM,
		cmd.Client.ID,
		"",
	)
	// 레디스 메모리에서 방 정보 삭제
	if err := ch.deleteRoomFromRedis(roomInfo.ID); err != nil {
		content := fmt.Sprintf("Failed to delete the room : %s", err.Error())
		ch.writeLogAndSendError(cmd.Client.ID, content)
	}

	log.Printf("User %s has left the room %s\n", cmd.Client.ID, roomInfo.ID)

	// 유저가 나갔다는 메시지 전달할 대상 탐색
	hostMap := ch.findBroadcastTargets(cmd, roomInfo)
	// 방 안의 다른 클라이언트에게 유저가 나갔다는 메시지 전송
	for hostId, clientIds := range hostMap {
		go ch.broadcastMessage(hostId,
			model.BROADCAST_OPPONENT_LEAVE_ROOM,
			cmd.Client.ID,
			clientIds,
			fmt.Sprintf("User %s has left the room.\n", cmd.Client.ID),
		)
	}

}

func (ch *ClientHandler) handleCreateRoom(cmd Command) {
	ch.Server.mutex.Lock()
	defer ch.Server.mutex.Unlock()

	// 유저가 방이 있는지 확인.
	if ch.checkClientHasRoom(cmd) {
		ch.writeLogAndSendError(cmd.Client.ID, "Client is already in a room")
		return
	}

	// 없는 고유햔 id 찾기
	roomId, err := ch.getNewRoomId(cmd)
	if err != nil {
		ch.writeLogAndSendError(cmd.Client.ID, err.Error())
		return
	}

	clients := make([]model.ClientInfo, 0, 1)
	clientInfo := model.ClientInfo{
		ID:     cmd.Client.ID,
		HostID: ch.Server.HostId,
	}
	clients = append(clients, clientInfo)
	roomInfo := &model.RoomInfo{
		ID:      roomId,
		Clients: clients,
	}

	if err := ch.setRoomInRedis(roomInfo); err != nil {
		ch.writeLogAndSendError(cmd.Client.ID, fmt.Sprintf("failed to create the room : %s", err.Error()))
		return
	}

	// 열려있는 방목록에 룸 정보 삽입
	err = ch.Server.RedisClient.RPush(ch.Server.ctx, "open_rooms", roomInfo.ToJson()).Err()
	if err != nil {
		content := fmt.Sprintf("Failed to create the room: %s", err.Error())
		ch.writeLogAndSendError(cmd.Client.ID, content)
		return
	}

	// 클라이언트 현재 방 업데이트
	cmd.Client.CurrentRoomId = roomId

	ch.sendMessageToClient(
		cmd.Client.ID,
		model.CLIENT_CREATE_ROOM_CONFIRM,
		cmd.Client.ID,
		"",
	)
	log.Printf("Client opened the new room %s\n", roomId)

}

func (ch *ClientHandler) handleSendMessageToOpponent(cmd Command) {

	if !ch.checkClientHasRoom(cmd) {
		ch.writeLogAndSendError(cmd.Client.ID, fmt.Sprintf("User %s does not belong to any room.", cmd.Client.ID))
		return
	}

	roomInfo, err := ch.getRoomInfo(cmd.Client.CurrentRoomId)

	if err != nil {
		ch.writeLogAndSendError(cmd.Client.ID, fmt.Sprintf("Failed to get the data from the redis. : %s", err.Error()))
		return
	}

	for _, c := range roomInfo.Clients {
		if c.ID == cmd.Client.ID {
			continue
		}

		msg := strings.Join(cmd.Args, " ") + "\n"
		ch.broadcastMessage(c.HostID,
			model.BROADCAST_OPPONENT_SEND_MESSAGE,
			cmd.Client.ID,
			[]string{c.ID},
			msg,
		)

	}

}

func (ch *ClientHandler) handleRemoveClient(cmd Command) {
	ch.Server.mutex.Lock()
	defer ch.Server.mutex.Unlock()

	if !ch.checkClientHasRoom(cmd) {
		delete(ch.Server.Clients, cmd.Client.ID)
		log.Printf("Removed client %s", cmd.Client.ID)
		return
	}

	roomInfo, err := ch.getRoomInfo(cmd.Client.CurrentRoomId)

	if err == nil {
		for _, client := range roomInfo.Clients {
			go ch.broadcastMessage(client.HostID, model.BROADCAST_OPPONENT_LEAVE_ROOM, cmd.Client.ID, []string{client.ID}, fmt.Sprintf("User %s has left the room.\n", cmd.Client.ID))
		}
	} else {
		log.Printf("Failed to get room info %s", roomInfo.ID)
	}

	err = ch.Server.RedisClient.Del(ch.Server.ctx, fmt.Sprintf("room:%s", cmd.Client.CurrentRoomId)).Err()
	if err != nil {
		log.Printf("Failed to delete the room that the client left. :%s\n", err.Error())
	}
	delete(ch.Server.Clients, cmd.Client.ID)
}

func (ch *ClientHandler) handleQuit(cmd Command) {

	ch.Server.mutex.Lock()
	defer ch.Server.mutex.Unlock()

	if !ch.checkClientHasRoom(cmd) {
		delete(ch.Server.Clients, cmd.Client.ID)
		cmd.Client.Conn.Close()
		log.Printf("Removed client %s", cmd.Client.ID)
		return
	}

	roomInfo, err := ch.getRoomInfo(cmd.Client.CurrentRoomId)

	if err == nil {
		for _, client := range roomInfo.Clients {
			go ch.broadcastMessage(client.HostID, model.BROADCAST_OPPONENT_LEAVE_ROOM, cmd.Client.ID, []string{client.ID}, fmt.Sprintf("User %s has left the room.\n", cmd.Client.ID))
		}
	} else {
		log.Printf("Failed to get room info %s", roomInfo.ID)
	}

	err = ch.Server.RedisClient.Del(ch.Server.ctx, fmt.Sprintf("room:%s", cmd.Client.CurrentRoomId)).Err()
	if err != nil {
		log.Printf("Failed to delete the room that the client left. :%s\n", err.Error())
	}
	delete(ch.Server.Clients, cmd.Client.ID)
	cmd.Client.Conn.Close()
}

// 다른 서버에 메시지 전달
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

// 이 서버랑 연결된 클라이언트에게 메시지 보내기
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

func (ch *ClientHandler) checkClientHasRoom(cmd Command) bool {
	if cmd.Client.CurrentRoomId == "" {
		return false
	} else {
		return true

	}
}

func (ch *ClientHandler) deleteRoomFromRedis(roomId string) error {
	err := ch.Server.RedisClient.Del(ch.Server.ctx, fmt.Sprintf("room:%s", roomId)).Err()
	if err != nil {
		return fmt.Errorf("failed to delete the room: %s", err.Error())
	}
	return nil
}

func (ch *ClientHandler) setRoomInRedis(roomInfo *model.RoomInfo) error {
	err := ch.Server.RedisClient.Set(ch.Server.ctx, fmt.Sprintf("room:%s", roomInfo.ID), roomInfo.ToJson(), time.Minute*30).Err()
	if err != nil {
		return err
	}
	return nil
}

func (ch *ClientHandler) findBroadcastTargets(cmd Command, roomInfo *model.RoomInfo) map[string][]string {
	hostMap := make(map[string][]string)
	for _, client := range roomInfo.Clients {
		// 나간 당사자면 메세지 전달 목록에 포함 안함
		if client.ID == cmd.Client.ID {
			continue
		}
		hostMap[client.HostID] = append(hostMap[client.HostID], client.ID)
	}
	return hostMap

}

func (ch *ClientHandler) getNewRoomId(cmd Command) (string, error) {
	var roomId string
	for {
		roomId = uuid.New().String()
		// 방정보 조회
		_, err := ch.getRoomInfo(roomId)
		if err == redis.Nil {
			break
		} else if err != nil {
			return "", fmt.Errorf("failed to create the room : %s", err.Error())
		} else {
			continue
		}

	}
	return roomId, nil
}

func (ch *ClientHandler) writeLogAndSendError(target string, content string) {
	log.Printf(content + "\n")
	ch.sendMessageToClient(target, model.CLIENT_ERROR, target, content)

}
