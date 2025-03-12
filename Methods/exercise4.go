package main

import "fmt"

type BankAccount struct {
	Balance float64
}

func (b *BankAccount) Deposit(amount float64) *BankAccount {
	b.Balance += amount
	return b
}
func (b *BankAccount) Withdraw(amount float64) *BankAccount {
	if amount <= b.Balance {
		b.Balance -= amount
	} else {
		fmt.Println("insufficient fund ")
	}
	return b

}
func main() {
	acc := BankAccount{}
	acc.Deposit(1000).Withdraw(10)
	fmt.Println("Final Balance ", acc.Balance)
}
