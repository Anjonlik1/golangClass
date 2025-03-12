package main

import (
	"fmt"
	"sync"
)

// TicketSystem encapsulates the total tickets and mutex
type TicketSale struct {
	totalTickets int
	mutex        sync.Mutex
}

// bookTicket is a method of TicketSystem
func (ts *TicketSale) bookTicket(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	ts.mutex.Lock()
	defer ts.mutex.Unlock()

	if ts.totalTickets > 0 {
		ts.totalTickets--
		fmt.Printf("Goroutine %d booked a ticket. Remaining tickets: %d\n", id, ts.totalTickets)
	} else {
		fmt.Printf("Goroutine %d tried to book a ticket, but none are left.\n", id)
	}
}

func main() {
	var wg sync.WaitGroup
	numGoroutines := 100

	// Initialize TicketSystem
	ts := TicketSale{
		totalTickets: 200,
	}

	wg.Add(numGoroutines)
	for i := 1; i <= numGoroutines; i++ {
		go ts.bookTicket(i, &wg)
	}

	wg.Wait()
	fmt.Printf("\nFinal remaining tickets: %d\n", ts.totalTickets)
}
