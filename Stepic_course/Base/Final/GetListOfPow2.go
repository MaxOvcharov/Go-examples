package main

import (
	"fmt"
)

func main() {
	var a uint
	fmt.Scan(&a)
	for i := uint(1); i <= a; i <<= 1 {
		fmt.Printf("%d ", i)
	}
}
