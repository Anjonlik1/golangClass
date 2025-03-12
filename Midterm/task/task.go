package task

import (
	"fmt"
)

type Product struct {
	Name     string
	Quantity int
	Price    float64
}

var inventory []Product

func AddProduct() {
	var Name string
	var Quantity int
	var Price float64

	fmt.Print("Enter product name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter prduct quantity: ")
	for {
		_, err := fmt.Scanln(&quantity)
		if err == nil && quantity >= 0 {
			break
		}
		fmt.Print("Invalid input")
	}

	fmt.Print("Enter the prodcut price: ")
	for {
		_, err := fmt.Scanln(&price)
		if err == nil && price >= 0 {
			break
		}
		fmt.Print("Invalid input.")
	}

	inventory = append(inventory, Product{Name: name, Quantity: quantity, Price: price})
	fmt.Println("Product added successfully!")

}

func ViewInventory() {
	fmt.Println("inventory:")
	for v, product := range inventory {
		fmt.Printf("%d. Name: %s, Quantity: %d, Price: $%.2f\n", v+1, product.Name, product.Quantity, product.Price)
	}

}
