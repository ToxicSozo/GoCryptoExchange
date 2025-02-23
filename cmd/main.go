package main

import (
	"log"

	"github.com/ToxicSozo/GoCryptoExchange/internal/app"
)

func main() {
	server := app.NewServer(":8080")

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
