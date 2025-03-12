package main

import "fmt"

type Celsius float64

func (t Celsius) ToFahrenheit() float64 {

	return float64(t)*9/5 + 32
}
func main() {
	var temp Celsius = 25
	fmt.Println(temp.ToFahrenheit())
}
