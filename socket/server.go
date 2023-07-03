package socket

import (
	"fmt"
	"net"
)

func StartServer() (net.Listener, error) {
	fmt.Println("Starting server...")

	listen, err := net.Listen("tcp", ":8000")
	if err != nil {
		return nil, err
	}

	return listen, err
}
