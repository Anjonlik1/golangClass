package main

import "fmt"

type Counter struct {
	value int
}

func (n *Counter) Increment() {
	n.value++

}

func (n *Counter) Value() int {

	return n.value
}
func main() {
	counter := Counter{value: 0}

	counter.Increment()
	counter.Increment()

	fmt.Println("Counter Value:", counter.Value())
	counter.Increment()
	fmt.Println("Update value: ", counter.Value())
}
