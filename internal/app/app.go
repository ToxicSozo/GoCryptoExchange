package app

import (
	"log"
	"net"

	"github.com/ToxicSozo/GoCryptoExchange/internal/handler"
)

type Server struct {
	addr string
}

func NewServer(addr string) *Server {
	return &Server{
		addr: addr,
	}
}

func (s *Server) Run() error {
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatal("Falied to start server", err)
		return err
	}

	defer ln.Close()

	log.Println("server is listening on :", s.addr)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("Error accept connection!")

			continue
		}

		go handler.HandlerConnection(conn)
	}
}
