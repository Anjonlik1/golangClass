package main

type BankAccount struct {
	Balance float64
}

func (b *BankAccount) Deposit(amount float64) *BankAccount {
	b.Balance += amount
	return b
}

func main() {

}
