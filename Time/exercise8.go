package main

import (
	"fmt"
	"time"
)

func main() {
	ts := int64(1707740400) // Example timestamp

	t := time.Unix(ts, 0) // Convert seconds to time.Time
	fmt.Println("Time from Unix:", t)

}
