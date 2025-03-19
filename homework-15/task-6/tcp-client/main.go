package main

import (
	"fmt"
	"net"
	"os"
)

/*
1. Make TCP client connection to server
2. Read data.txt file from the file system
3. Send file to the server
*/

func main() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Printf("Error TCP client connection: %s\n", err)
	}
	fmt.Println("Connected to TCP server")
	fileBytes, err := ReadFile("data.txt")
	if err != nil {
		fmt.Printf("Error read file: %s\n", err)
	}
	_, err = conn.Write(fileBytes)
	if err != nil {
		fmt.Printf("Error send file: %s\n", err)
	}
	conn.Close()
}

func ReadFile(fileName string) ([]byte, error) {
	return os.ReadFile(fileName)
}
