package main

import "fmt"

type Human struct {
	Name     string
	Greeting string
}

func (h Human) Greet() {
	fmt.Printf("%s says %s", h.Name, h.Greeting)
}
func main() {
	odam := Human{
		Name:     "Andy",
		Greeting: "Assalam Alaykom",
	}
	odam.Greet()
}
