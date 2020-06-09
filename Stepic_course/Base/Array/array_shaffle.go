package main

import (
	"fmt"
)

func main() {
	var a, b int
	var workArray [10]uint8

	for i := 0; i < len(workArray); i++ {
		fmt.Scan(&workArray[i])
	}

	for i := 0; i < 3; i++ {
		fmt.Scan(&a, &b)
		workArray[b], workArray[a] = workArray[a], workArray[b]
	}

	for _, v := range workArray {
		fmt.Printf("%v ", int(v))
	}
}