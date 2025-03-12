package main

import "fmt"

func main() {
	myChannel := make(chan int)

	go func() {
		myChannel <- 42
	}()
	value := <-myChannel
	fmt.Println("recieved value: ", value)
}
