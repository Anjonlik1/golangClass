package inventory

import (
	"fmt"
)

type Product struct {
	Name     string
	Quantity int
	Price    float64
}

var Inventory []Product

func AddProduct() {
	var Name string
	var Quantity int
	var Price float64

	fmt.Print("Enter product name: ")
	fmt.Scanln(&Name)

	fmt.Print("Enter prduct quantity: ")
	for {
		_, err := fmt.Scanln(&Quantity)
		if err == nil && Quantity >= 0 {
			break
		}
		fmt.Print("Invalid input")
	}

	fmt.Print("Enter the prodcut price: ")
	for {
		_, err := fmt.Scanln(&Price)
		if err == nil && Price >= 0 {
			break
		}
		fmt.Print("Invalid input.")
	}

	inventory = append(Inventory, Product{Name: Name, Quantity: Quantity, Price: Price})
	fmt.Println("Product added successfully!")

}

func ViewInventory() {
	fmt.Println("inventory:")
	for v, product := range Inventory {
		fmt.Printf("%d. Name: %s, Quantity: %d, Price: $%.2f\n", v+1, product.Name, product.Quantity, product.Price)
	}

}
