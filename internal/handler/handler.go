package handler

import (
	"bufio"
	"fmt"
	"net"
)

func HandlerConnection(conn net.Conn) {
	defer conn.Close()
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Print("Message Received:", string(message))
		conn.Write([]byte("Message received: " + message))
	}
}
