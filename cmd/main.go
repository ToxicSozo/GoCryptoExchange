package main

import (
	"log"

	"github.com/ToxicSozo/GoCryptoExchange/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
