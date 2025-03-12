package main

import "fmt"

type Rectangle struct {
	Width  int
	Height int
}

func (r Rectangle) Area() int {
	return r.Height * r.Width

}
func (r Rectangle) Perimeter() int {
	return (r.Height + r.Width) * 2

}

func main() {
	r := Rectangle{Width: 3, Height: 4}
	fmt.Println(r.Area())
	fmt.Println(r.Perimeter())
}
