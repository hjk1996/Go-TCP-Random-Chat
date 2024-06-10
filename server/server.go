package main

import "log"

type Server struct {
	ComChan chan Command
	Rooms   map[string]*Room
}

func NewServer() *Server {
	return &Server{
		ComChan: make(chan Command),
		Rooms:   make(map[string]*Room),
	}
}

func (server *Server) Run() {
	for cmd := range server.ComChan {
		switch cmd.CommandType {
		case CMD_JOIN_ROOM:
			server.JoinRoom(cmd)
		case CMD_LEAVE_ROOM:
			server.LeaveRoom(cmd)
		case CMD_NEW_ROOM:
			server.CreateRoom(cmd)
		case CMD_CHANGE_NICK:
			server.ChangeNick(cmd)
		case CMD_SEND_MESSAGE:
			server.SendMessage(cmd)

		}

	}
}

func (server *Server) JoinRoom(cmd Command) {
	log.Println("join")
}

func (server *Server) LeaveRoom(cmd Command) {
	log.Println("leave")

}

func (server *Server) CreateRoom(cmd Command) {
	log.Println("create")

}

func (server *Server) ChangeNick(cmd Command) {
	log.Println("change")

}

func (server *Server) SendMessage(cmd Command) {
	log.Println("msg")

}
