package main

import (
	"fmt"
	"net"
	"os"
)

/*
	1. Create TCP server connection
	2. Accept TCP client connections
	3. Parse data from client request
	4. Save file in server side
	5. Notify client side
*/

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("Error create TCP server: %s\n", err)
	}
	fmt.Println("TCP server started successfully")
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Error accepting client connection: %s\n", err)
		}
		go HandleIncomingRequest(conn)
	}
}

func HandleIncomingRequest(conn net.Conn) {
	response := ""
	fileBuf := make([]byte, 1024) // 1024
	clientAddr := conn.RemoteAddr().String()

	fmt.Printf("Client with address %s connected\n", clientAddr)
	n, err := conn.Read(fileBuf) // fileBuf
	fmt.Println(string(fileBuf))
	if err != nil {
		fmt.Printf("Error read data: %s\n", err)
		response = err.Error()
	}
	defer conn.Close()
	fileName := fmt.Sprintf("data(%s).txt", clientAddr) // data(127.0.0.1:6545).txt
	err = SaveFile(fileBuf[:n], fileName)
	if err != nil {
		fmt.Printf("Error save file: %s\n", err)
		response = err.Error()
	}

	response = "File successfully saved!"
	conn.Write([]byte(response)) // Send to the client side
	return
}

func SaveFile(fileBytes []byte, fileName string) error {
	return os.WriteFile(fileName, fileBytes, 0666)
}
