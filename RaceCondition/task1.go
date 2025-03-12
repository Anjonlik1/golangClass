package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func BankBalance(balance *int64, amount int64, wg *sync.WaitGroup) {
	defer wg.Done()
	atomic.AddInt64(balance, amount)
}
func main() {
	var balance int64 = 0
	var wg sync.WaitGroup

	numTransaction := 5
	amountToAdd := int64(100)

	wg.Add(numTransaction)

	for i := 0; i < numTransaction; i++ {
		go BankBalance(&balance, amountToAdd, &wg)
	}

	wg.Wait()

	fmt.Printf("Final Balance: %d\n", balance)

}
