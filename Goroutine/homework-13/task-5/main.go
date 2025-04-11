package main

import "fmt"

func main() {
	input := make(chan int, 10)
	output := make(chan int, 10)

	go generator(input)
	go doubler(input, output)

	printer(output)

}

func generator(input chan int) {
	for i := 1; i <= 10; i++ {
		input <- i
	}
	close(input)
}

func doubler(input, output chan int) {
	for v := range input {
		output <- 2 * v
	}
	close(output)
}

func printer(output chan int) {
	for v := range output {
		fmt.Println(v)
	}
}
