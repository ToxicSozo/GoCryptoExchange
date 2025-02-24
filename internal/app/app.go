package app

import (
	"log"
	"net"

	h "github.com/ToxicSozo/GoCryptoExchange/internal/handler"
)

const (
	HOST = "localhost"
	PORT = ":8080"
	TYPE = "tcp"
)

func Run() error {
	listen, err := net.Listen(TYPE, HOST+PORT)
	if err != nil {
		return err
	}

	defer listen.Close()

	log.Println("server is listening on ", PORT)

	for {
		conn, err := listen.Accept()
		if err != nil {
			return err
		}

		go h.HandleRequest(conn)
	}
}
