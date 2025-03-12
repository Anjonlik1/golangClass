package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func randomNumber(numbers []int, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	numbers[i] = rand.Intn(100)
}
func main() {
	n := 5
	numbers := make([]int, n)
	var wg sync.WaitGroup

	wg.Add(n)

	for i := 0; i < n; i++ {
		go randomNumber(numbers, i, &wg)
	}

	wg.Wait()
	fmt.Println("Generated numbers:", numbers)
}
