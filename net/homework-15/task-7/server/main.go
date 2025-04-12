package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

var response string

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Error listening on port 8080: %s\n", err)
	}
	defer listener.Close()
	fmt.Printf("Listening TCP server on port 8080\n")

	conn, err := listener.Accept() // Client
	if err != nil {
		fmt.Printf("Error accepting connection: %s\n", err)
	}
	handleConnection(conn)

}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	clientAddress := conn.RemoteAddr().String()
	fmt.Printf("Connection with client %v has established\n", clientAddress)
	reader := bufio.NewReader(conn)

infinite:
	for {
		command, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Client %s disconnected\n", clientAddress)
			break
		}
		switch {
		case strings.HasPrefix(command, "ECHO:"):
			arr := strings.Split(command, ":")
			if len(arr) > 2 {
				response = "Invalid command\n"
				conn.Write([]byte(response))
				continue
			}
			response = arr[1] + "\n"
		case command == "TIME\n":
			response = time.Now().Format(time.Stamp) + "\n"

		case command == "EXIT\n":
			conn.Write([]byte("Shutting down..."))
			break infinite
		default:
			response = "Invalid command\n"
		}
		conn.Write([]byte(response))
	}
}
