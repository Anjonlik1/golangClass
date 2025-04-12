package main

import (
	"bufio"
	"fmt"
	"net"
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

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Client %s disconnected\n", clientAddress)
			break
		}
		res := reverseString(message)
		fmt.Println(res)
		conn.Write([]byte(res + "\n"))
	}
}

func reverseString(s string) string {
	var res string
	for i := len(s) - 1; i >= 0; i-- {
		res = res + string(s[i])
	}
	return res
}
