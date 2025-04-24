package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	count := 0
	for range ticker.C {
		fmt.Println("Task Executed")

		count++
		if count >= 10 {
			break
		}
	}
}
