package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var photoURL = "https://youtu.be/WqrbIUggEXQ?si=EbR_Py8fgPcRaCZ5"

/*

	1. Create HTTP client
	2. Send GET request via photoURL
	3. Parse body from server response
	4. Save image in file system (local computer)

*/

func main() {
	// Step 1
	client := http.Client{}

	// Step 2
	response, err := client.Get(photoURL)
	if err != nil {
		fmt.Printf("Error make request: %s\n", err)
	}
	defer response.Body.Close()

	// Step 3
	fileBytes, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error parse body: %s\n", err)
	}

	// Step 4
	err = SaveImage(fileBytes, "placeholder.mp4")
	if err != nil {
		fmt.Printf("Error save image: %s\n", err)
	}

}

func SaveImage(image []byte, fileName string) error {
	return os.WriteFile(fileName, image, 0666)
}
