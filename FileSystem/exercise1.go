package main

import (
	"log"
	"os"
)

func createFile() {
	filePtr, err := os.Create("example.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer filePtr.Close()
}
func writeFileContent() {
	content := "Hello, Golang!"
	err := os.WriteFile("example.txt", []byte(content), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	createFile()
	writeFileContent()

}
