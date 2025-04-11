package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan string)
	timer := time.After(3 * time.Second)

	go func() {
		time.Sleep(2 * time.Second)
		done <- "done"
	}()

	select {
	case <-timer:
		fmt.Println("timeout") // Deadline
	case v := <-done:
		fmt.Println(v)
	}
}
