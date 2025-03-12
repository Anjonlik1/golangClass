package main

import (
	"bufio"
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	clientAddress := conn.RemoteAddr().String()
	fmt.Printf("New client connected to TCP server from %v\n", clientAddress)
	reader := bufio.NewReader(conn)
	for {
		msg, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Client %s disconnected\n", clientAddress)
			break
		}
		fmt.Printf("[%v] > %s\n", clientAddress, msg)
		response := "Hello from TCP server!\n"
		conn.Write([]byte(response))
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Error listening on port 8080: %s\n", err)
	}
	defer listener.Close()
	fmt.Printf("Listening TCP server on port 8080\n")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting connection: %s\n", err)
		}
		go handleConnection(conn)
	}
}
