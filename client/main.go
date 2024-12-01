package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// соединяемся с сервером
	conn, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		fmt.Println("Error connecting:", err.Error())
		os.Exit(1)
	}
	defer conn.Close()

	// читаем сообщения с консоли и отправляем их серверу
	consoleScanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter text to send:")
	for consoleScanner.Scan() {
		text := consoleScanner.Text()
		conn.Write([]byte(text + "\n"))

		// получаем ответ от сервера
		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			os.Exit(1)
		}
		fmt.Print("Server says: " + response)
		fmt.Println("Enter more text to send:")
	}

	if err := consoleScanner.Err(); err != nil {
		fmt.Println("Error reading from console:", err.Error())
	}
}
