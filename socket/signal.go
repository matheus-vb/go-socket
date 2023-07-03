package socket

import (
	"fmt"
	"log"
	"net"
	"sync"
)

func Signal(clients *[]*net.Conn, message string, lock *sync.Mutex) error {
	lock.Lock()
	defer lock.Unlock()

	for _, connection := range *clients {
		_, err := (*connection).Write([]byte(message))
		if err != nil {
			fmt.Println("Error writing message")
			return err
		}
		log.Println("Sent signal")
	}

	fmt.Println("Signal done")
	return nil
}
