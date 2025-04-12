package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", ":8080") // TCP server
	if err != nil {
		fmt.Printf("Error connecting to TCP server: %v\n", err)
	}
	defer conn.Close()
	fmt.Printf("Connection was established\n")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("You > ")
		message, _ := reader.ReadString('\n')

		_, err := conn.Write([]byte(message))
		if err != nil {
			fmt.Printf("Error writing to TCP server: %v\n", err)
		}
		serverResponse, _ := bufio.NewReader(conn).ReadString('\n')

		fmt.Println("Server > ", serverResponse)
		if strings.HasPrefix(serverResponse, "Shutting down") {
			break
		}
	}
}
