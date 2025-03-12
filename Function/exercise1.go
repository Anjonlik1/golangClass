package main

import "fmt"

type Product struct {
	Name  string
	Price float64
}

func totalPrice(products ...Product) float64 {
	total := 0.0
	for _, product := range products {
		total += product.Price
	}
	return total
}

func main() {
	p1 := Product{Name: "My wife dress", Price: 9999}
	p2 := Product{Name: "O'glimi legosi", Price: 500}

	fmt.Println("Total price: ", totalPrice(p1, p2))

}
