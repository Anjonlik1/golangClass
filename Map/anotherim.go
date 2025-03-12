package main

import "fmt"

type Employee struct {
	Name       string
	Department string
}

func main() {

	employeeDirectory := map[int]Employee{
		1: {"Toshmirza", "Research"},
		2: {"Eshmirza", "Human Resource"},
		3: {"Toshbolta", "Production"},
		4: {"Sotvoldi", "Sales"},
	}

	fmt.Println("Employee Directory:")
	for id, emp := range employeeDirectory {
		fmt.Printf("ID: %d, Name: %s, Department: %s\n", id, emp.Name, emp.Department)
	}
	fmt.Println("-------------------------------")
	employeeDirectory[5] = Employee{"Boltavoy", "Ads"}

	fmt.Println("New Employee Directory:")
	for id, emp := range employeeDirectory {
		fmt.Printf("ID: %d, Name: %s, Department: %s\n", id, emp.Name, emp.Department)
	}

	fmt.Println("------------")
	delete(employeeDirectory, 2)
	fmt.Println("After Deletion Directory:")
	for id, emp := range employeeDirectory {
		fmt.Printf("ID: %d, Name: %s, Department: %s\n", id, emp.Name, emp.Department)
	}
}
