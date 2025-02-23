package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080") // Исправлен порт
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close() // Закрываем соединение при завершении

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Text to send: ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading input:", err)
			break
		}

		_, err = conn.Write([]byte(text)) // Отправляем без лишнего \n
		if err != nil {
			log.Println("Error writing to server:", err)
			break
		}

		message, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Println("Error reading from server:", err)
			break
		}
		fmt.Print("Message from server: " + message)
	}
}
