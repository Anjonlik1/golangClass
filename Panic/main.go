package main

import (
	"fmt"
)

func accessElement(arr []int, index int) (int, error) {
	if index < 0 || index >= len(arr) {
		panic(fmt.Sprintf("Index %d out of bounds!", index))
	}
	// return arr[index], nil
	 fmt.Println("Element:", arr[index]), nil
}

func main() {
	arr := []int{10, 20, 30}
	accessElement(arr, 5) // This will cause a panic

	element, err := accessElement(arr, 5)
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("element:", element)
	}

	element, err = accessElement(arr, 1)
	if err != nil {
		fmt.Println("error: ", err)
	} else {
		fmt.Println("element:", element)
	}

}
