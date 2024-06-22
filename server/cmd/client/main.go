package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"example.com/chat/pkg/client"
)

func connectLoop() *client.Client {

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter server address: ")
		addr, _ := reader.ReadString('\n')
		client := client.NewClient(addr)

		err := client.Connect()

		if err != nil {
			log.Println(err.Error())
		}

		return client
	}

}

func main() {

}
