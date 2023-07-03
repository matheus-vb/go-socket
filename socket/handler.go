package socket

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"sync"
)

func Handle(connection *net.Conn, clients *[]*net.Conn, lock *sync.Mutex) error {
	for {
		message, err := bufio.NewReader(*connection).ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("Closed connection")
				(*connection).Close()

				lock.Lock()
				for i, client := range *clients {
					if *client == *connection {
						*clients = append((*clients)[:i], (*clients)[i+1:]...)
						break
					}
				}
				lock.Unlock()
			}

			fmt.Println("Error reading message")
			return err
		}

		_, err = (*connection).Write([]byte(message))
		if err != nil {
			fmt.Println("Error writing message")
			return err
		}

		go Signal(clients, message, lock)

		fmt.Printf("%v, %q", (*connection).RemoteAddr(), message)
	}
}
