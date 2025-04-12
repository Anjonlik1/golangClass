package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
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
		serverResponse, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("Connection was closed")
			break
		}
		fmt.Println("Result > ", serverResponse)
	}
}
