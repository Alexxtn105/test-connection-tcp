package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// определяем порт для прослушивания
	PORT := ":9090"
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// закрываем listener при завершении программы
	defer listener.Close()
	fmt.Println("Server is listening on " + PORT)

	for {
		// принимаем входящее подключение
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting:", err.Error())
			os.Exit(1)
		}
		fmt.Println("Connected with", conn.RemoteAddr().String())
		// обрабатываем подключение в отдельной горутине
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	// читаем данные от клиента
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		clientMessage := scanner.Text()
		fmt.Printf("Received from client: %s\n", clientMessage)
		// отправляем ответ клиенту
		conn.Write([]byte("Message received.\n"))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading:", err.Error())
	}
}
