package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	// Read input from user
	fmt.Print("Enter a string: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Send the string to the server
	fmt.Fprintln(conn, input)

	// Receive and print the reversed string
	response, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("Server is sending Reversed string:", strings.TrimSpace(response))
}
