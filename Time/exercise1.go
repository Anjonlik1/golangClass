package main

import (
	"fmt"

	"time"
)

func main() {
	timestamp := "2025-02-12 13:15:35"
	ts, err := time.Parse("2006-01-02 15:04:05", timestamp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ts)

	format1 := ts.Format("2006-01-02")
	fmt.Println(format1)

	format2 := ts.Format("02/01/2006")
	fmt.Println(format2)

	format3 := ts.Format("Monday, 02-Jan-2006")
	fmt.Println(format3)
}
