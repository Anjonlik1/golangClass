package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	t := time.Now()
	zone, offset := t.Zone()
	fmt.Println("local zone:", zone)
	fmt.Println("time offset,", offset)
}
