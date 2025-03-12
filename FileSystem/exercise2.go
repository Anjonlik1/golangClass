package main

import (
	"fmt"
	"log"
	"os"
)

func createFile() {
	filePtr, err := os.Create("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer filePtr.Close()
}
func writeFileContent() {
	content := "Hello, Golang!"
	err := os.WriteFile("sample.txt", []byte(content), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
func readFileContents() {
	bytes, err := os.ReadFile("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	fileText := string(bytes)
	fmt.Println(fileText)
}
func main() {
	createFile()
	writeFileContent()
	readFileContents()

}
