package main

import "fmt"

type Employee struct {
	Name       string
	Department string
}

func main() {
	employeeDirectory := make(map[int]Employee)

	employeeDirectory[1] = Employee{Name: "Toshmirza", Department: "Research"}
	employeeDirectory[2] = Employee{Name: "Eshmirza", Department: "Human Resource"}
	employeeDirectory[3] = Employee{Name: "Toshbolta", Department: "Production"}
	employeeDirectory[4] = Employee{Name: "Sotvoldi", Department: "Sales"}

	fmt.Println("Employee Directory : ")

	fmt.Println("ID\tName\t\tDepartment")
	fmt.Println()

	for empID, employee := range employeeDirectory {
		fmt.Printf("%d\t%s\t%s\n", empID, employee.Name, employee.Department)
	}

	delete(employeeDirectory, "Toshbolta")

	fmt.Println(employeeDirectory)

}
