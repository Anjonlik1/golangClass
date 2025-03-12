package main

import "fmt"

/*
### **Buffered Channel**

- Create a **buffered channel** with a capacity of 2.
- Send two numbers (`10` and `20`) into the channel.
- Receive both numbers and print them.
*/
func main() {

	myChannel := make(chan int, 2)
	go func() {
		myChannel <- 10
	}()

	go func() {
		myChannel <- 20
	}()

	val1 := <-myChannel
	val2 := <-myChannel
	fmt.Println("recieved first value:", val1)
	fmt.Println("recieved second value:", val2)
	// git process
}
