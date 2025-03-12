/*
1. **Custom Error for Negative Numbers**
  - Create a custom error type `NegativeNumberError` with a field `Value`.
  - Implement the `Error()` method to return `"negative number encountered: Value"`.
  - Write a function `sqrt(n float64) (float64, error)`.
  - If `n < 0`, return `NegativeNumberError`.
  - Handle the error in `main()`
*/
package main

import (
	"fmt"
	"math"
)

type NegativeNumberError struct {
	Value float64
}

func (n NegativeNumberError) Error() string {
	return fmt.Sprintf("negative number: %v", n.Value)
}

func sqrt(n float64) (float64, error) {
	if n < 0 {
		return 0, NegativeNumberError{Value: n}
	}
	return math.Sqrt(n), nil
}

func main() {
	result, err := sqrt(10)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("square root:", result)
	}

	result, err = sqrt(-4)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("square root:", result)
	}

}
