package main

import (
	"fmt"

	"github.com/Surya-7890/tcp-test/server"
)

func main() {
	new_server := server.NewServer(":7000")
	go func() {
		for msg := range new_server.Msg {
			fmt.Println("message: ", string(msg))
		}
	}()
	new_server.Start()
}
