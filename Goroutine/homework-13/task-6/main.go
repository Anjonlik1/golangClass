package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	ticker := time.NewTicker(2 * time.Second)
	done := make(chan bool)

	go func() {
		time.Sleep(17 * time.Second)
		done <- true
	}()

label:
	for {
		select {
		case t := <-ticker.C:
			fmt.Printf("Task executed %v\n", t)
		case <-done:
			ticker.Stop()
			break label
		}
	}

	fmt.Println(time.Since(start))

}

// count == 6
// :00

// count = 1 log //:02
// count = 2 log :04
// count = 3 log :06
// count = 4 log :08
// count = 5 log :10
// count = 6 log :12
// count = 7 log :14
// count = 8 log :16
// count = 9 log :18

// Start time = 16:20:20
// End time = 16:20:28
// Diff = 8s
