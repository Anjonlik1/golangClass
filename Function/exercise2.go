package main

import "fmt"

func filter(nums []int, predicate func(int) bool) []int {
	filterNumbers := []int{}

	for _, number := range nums {
		if predicate(number) {
			filterNumbers = append(filterNumbers, number)
		}
	}
	return filterNumbers
}
func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	evens := filter(nums, func(number int) bool {
		return number%2 == 0
	})
	fmt.Println(evens) // Output: [2 4 6]

}
