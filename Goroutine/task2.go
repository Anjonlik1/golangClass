package main

import (
	"fmt"
	"strings"
	"sync"
)

func toUpper(s []string, i int, wg *sync.WaitGroup) {
	defer wg.Done()
	s[i] = strings.ToUpper(s[i])
}

func main() {
	s := []string{"hello", "world", "golang"}

	var wg sync.WaitGroup
	wg.Add(len(s))
	for i := range s {

		go toUpper(s, i, &wg)

	}
	wg.Wait()
	fmt.Println(s)
}
