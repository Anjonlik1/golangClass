package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

func readFileContents(filename string, wg *sync.WaitGroup) {
	defer wg.Done()
	bytes, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("Error reading %s: %v\n", filename, err)
		return
	}
	fmt.Printf("Contents of %s:\n%s\n", filename, bytes)

}
func main() {
	files := []string{"file1.txt", "file2.txt", "file3.txt"}
	var wg sync.WaitGroup

	wg.Add(len(files))

	for _, file := range files {
		go readFileContents(file, &wg)
	}

	wg.Wait()
	fmt.Println("Reading completed.")
}
