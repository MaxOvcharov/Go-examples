package main

import "fmt"

func sum(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Printf("NUMS: %d = %d\n", nums, total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}