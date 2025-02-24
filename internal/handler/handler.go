package handler

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func HandleRequest(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))

		message, err := reader.ReadString('\n')
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				continue
			}

			fmt.Println("Клиент отключился")
			break
		}

		fmt.Printf("Получено сообщение: %s", message)
	}
}
