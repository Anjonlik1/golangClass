package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {

		time.Sleep(time.Second)
		results <- j * 2
		fmt.Printf("Worker %d processed job %d\n", id, j)

	}

}

func main() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)
	for w := 1; w <= 5; w++ {
		go worker(w, jobs, results)
	}
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a <= 5; a++ {
		result := <-results
		fmt.Printf("Result %d: %d\n", a, result)
	}
	fmt.Println("All jobs processed")

}
