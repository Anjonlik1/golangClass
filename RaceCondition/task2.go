package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var counter int64 = 0

func countGoroutine(wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt64(&counter, 1)
}
func main() {
	var wg sync.WaitGroup
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go countGoroutine(&wg)
	}

	wg.Wait()
	fmt.Println("Counter =", counter)
}
