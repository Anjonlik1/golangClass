package main

import "fmt"

type Student struct {
	Name    string
	Grades  []int
	Faculty string
	Id      int
}

func (s Student) Detail() string {
	return s.Name
}
func main() {
	fmt.Println("methods")
	s1 := Student{Name: "Sam", Grades: []int{5, 4, 5, 3, 5}}
	s1.Detail()
	fmt.Println(s1)
}
