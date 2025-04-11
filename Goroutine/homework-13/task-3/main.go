package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		jobs <- i
	}

	close(jobs)

	wg.Add(3)
	go func() {
		defer wg.Done()
		worker(1, jobs, results)
	}()
	go func() {
		defer wg.Done()
		worker(2, jobs, results)
	}()

	go func() {
		defer wg.Done()
		printer(results)
	}()
	wg.Wait()
}

func printer(results chan int) {
	for r := range results {
		fmt.Println("Result: ", r)
	}
}

func worker(id int, jobs <-chan int, results chan int) {

	for v := range jobs {
		fmt.Printf("Worker #%d is processing job %d\n", id, v)
		results <- v * 2
	}
}
