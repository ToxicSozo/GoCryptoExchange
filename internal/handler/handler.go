package handler

import (
	"bufio"
	"fmt"
	"net"
	"time"

	"github.com/ToxicSozo/GoCryptoExchange/internal/utils"
)

func HandleRequest(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		conn.SetReadDeadline(time.Now().Add(10 * time.Second))

		query, err := reader.ReadString('\n')
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				continue
			}

			fmt.Println("Клиент отключился")
			break
		}

		parsed, err := utils.ParseQuery(query)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Printf("Parsed Query: %+v\n", parsed)
		}
	}
}
