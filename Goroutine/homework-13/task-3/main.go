package main

import (
	"fmt"
)

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	go worker(1, jobs, results)
	go worker(2, jobs, results)
	go worker(3, jobs, results)

	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)

	printer(results)
}

func printer(results chan int) {
	for a := 1; a <= 5; a++ {
		fmt.Println("Result:", <-results)
	}
	close(results)
}

func worker(id int, jobs <-chan int, results chan int) {

	for v := range jobs {
		fmt.Printf("Worker #%d is processing job %d\n", id, v)
		results <- v * 2
	}
}
