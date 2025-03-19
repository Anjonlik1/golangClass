package main

import (
	"fmt"
	"sync"
)

func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	numbers := []int{5, 7, 10, 12}
	results := make(map[int]int)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fact := factorial(n)
			mu.Lock()
			results[n] = fact
			mu.Unlock()
		}(num)
	}

	wg.Wait()

	// Print results
	for num, fact := range results {
		fmt.Printf("Factorial of %d is %d\n", num, fact)
	}
}
