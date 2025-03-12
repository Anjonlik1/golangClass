// package main

// import (
// 	"fmt"
// )

// type Product struct {
// 	Name     string
// 	Quantity int
// 	Price    float64
// }

// var inventory []Product

// func AddProduct() {
// 	var name string
// 	var quantity int
// 	var price float64

// 	fmt.Print("Enter product name: ")
// 	fmt.Scanln(&name)

// 	fmt.Print("Enter prduct quantity: ")
// 	for {
// 		_, err := fmt.Scanln(&quantity)
// 		if err == nil && quantity >= 0 {
// 			break
// 		}
// 		fmt.Print("Invalid input")
// 	}

// 	fmt.Print("Enter the prodcut price: ")
// 	for {
// 		_, err := fmt.Scanln(&price)
// 		if err == nil && price >= 0 {
// 			break
// 		}
// 		fmt.Print("Invalid input.")
// 	}

// 	inventory = append(inventory, Product{Name: name, Quantity: quantity, Price: price})
// 	fmt.Println("Product added successfully!")

// }

// func ViewInventory() {
// 	fmt.Println("inventory:")
// 	for v, product := range inventory {
// 		fmt.Printf("%d. Name: %s, Quantity: %d, Price: $%.2f\n", v+1, product.Name, product.Quantity, product.Price)
// 	}

// }
// func main() {
// 	for {
// 		fmt.Println("\nMenu:")
// 		fmt.Println("1. Add Product")
// 		fmt.Println("2. View Products")
// 		fmt.Println("3. Exit")

// 		var option int
// 		fmt.Scanln(&option)

// 		switch option {
// 		case 1:
// 			AddProduct()
// 		case 2:
// 			ViewInventory()
// 		case 3:
// 			fmt.Println("Goodbye")
// 			return
// 		default:
// 			fmt.Println("Invalid option.")
// 		}
// 	}

// }
