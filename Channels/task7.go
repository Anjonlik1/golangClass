package main

import (
	"fmt"
	"time"
)

func main() {
	chan := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		chan <- "Done"

	}()
	select {
	case msg := <-done:
		fmt.Println("Received", msg)
	case <-time.After(3 * time.Second):
		fmt.Println("Timeout")
	}
}
