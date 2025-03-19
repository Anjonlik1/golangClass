package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Printf("Error connecting to TCP server: %v\n", err)
	}
	defer conn.Close()
	fmt.Printf("Connected to TCP server\n")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("You > ")
		message, _ := reader.ReadString('\n')

		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Printf("Error writing to TCP server: %v\n", err)
		}
		serverResponse, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading from TCP server: %v\n", err)
		}
		fmt.Println("Server > ", string(serverResponse))
	}
}
