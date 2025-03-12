package main

import (
	"fmt"
	"sync"
)

func print(n int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(n)

}

func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go print(i, &wg)
	}
	wg.Wait()

	fmt.Println("Printing finished")
}
