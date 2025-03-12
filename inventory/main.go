package main

import (
	"fmt"
)

func main() {
	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Add Product")
		fmt.Println("2. View Products")
		fmt.Println("3. Exit")

		var option int
		fmt.Scanln(&option)

		switch option {
		case 1:
			AddProduct()
		case 2:
			ViewInventory()
		case 3:
			fmt.Println("Goodbye")
			return
		default:
			fmt.Println("Invalid option.")
		}
	}

}
