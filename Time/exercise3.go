package main

import (
	"fmt"
	"time"
)

func main() {
	loc, _ := time.LoadLocation("America/New_York") // Load Eastern Time (ET)
	t := time.Now().In(loc)

	fmt.Println("Time in New York:", t)

	loc1, _ := time.LoadLocation("Asia/Tokyo") // Load Eastern Time (ET)
	t2 := time.Now().In(loc1)

	fmt.Println("Time in Tokyo:", t2)

	loc3, _ := time.LoadLocation("America/New_York") // Load Eastern Time (ET)
	t3 := time.Now().In(loc3)

	fmt.Println("Time in Morgantown:", t3)

}
