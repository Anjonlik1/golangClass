package main

import (
	"fmt"
)

type InsufficientFundsError struct {
	Balance float64
	Amount  float64
}

func (a *InsufficientFundsError) Error() string {
	return fmt.Sprintf("insufficient funds: balance = %.2f, required = %.2f", a.Balance, a.Amount)
}
func withdraw(balance, amount float64) (float64, error) {
	if amount > balance {
		return balance, &InsufficientFundsError{Balance: balance, Amount: amount}
	}
	return balance - amount, nil
}
func main() {
	balance := 50
	amount := 100

	newBalance, err := withdraw(float64(balance), float64(amount))

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Withdraw successful ! %.2f\n", newBalance)
	}
}
