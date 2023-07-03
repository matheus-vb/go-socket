package main

import (
	"fmt"
	"net"
	"sync"

	"github.com/matheus-vb/go-socket/socket"
)

func main() {
	listen, err := socket.StartServer()
	if err != nil {
		fmt.Println(err)
	}

	var clients []*net.Conn
	var lock sync.Mutex

	for {
		connection, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		fmt.Println("New connection")

		lock.Lock()
		clients = append(clients, &connection)
		lock.Unlock()

		go socket.Handle(&connection, &clients, &lock)
	}
}
