package server

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"

	"github.com/google/uuid"
)

type User struct {
	ID            string
	Conn          net.Conn
	ComChan       chan<- Command
	CurrentRoomId string
}

func NewUser(conn net.Conn, comChan chan<- Command) *User {
	clientId := uuid.New().String()
	return &User{
		ID:      clientId,
		Conn:    conn,
		ComChan: comChan,
	}
}

func (c *User) ReadInput() {
	for {
		msg, err := bufio.NewReader(c.Conn).ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Printf("Connection closed by client %s", c.ID)

				c.ComChan <- Command{
					Client:      c,
					CommandType: CMD_REMOVE_CLIENT,
				}
				return
			}
			log.Printf("Failed to read user input: %s", err.Error())
		}

		chunks := strings.Split(msg, " ")
		cmd := strings.TrimSpace(chunks[0])

		switch cmd {
		case "/join":
			c.ComChan <- Command{
				Client:      c,
				CommandType: CMD_JOIN_ROOM,
			}

		case "/leave":
			c.ComChan <- Command{
				Client:      c,
				CommandType: CMD_LEAVE_ROOM,
			}
		case "/new-room":
			c.ComChan <- Command{
				Client:      c,
				CommandType: CMD_NEW_ROOM,
			}
		case "/msg":
			if len(chunks) < 2 {
				log.Printf("wrong message command.\n")
				continue
			}
			c.ComChan <- Command{
				Client:      c,
				CommandType: CMD_SEND_MESSAGE,
				Args:        chunks[1:],
			}
		default:
			log.Printf("Wrong command. : %s\n", cmd)
		}

	}
}
