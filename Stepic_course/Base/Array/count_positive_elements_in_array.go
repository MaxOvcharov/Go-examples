package main

import (
	"fmt"
)

func main() {
	var index int

	fmt.Scan(&index)

	var intArray [100]int

	for i := 0; i != index; i++ {
		fmt.Scan(&intArray[i])
	}

	var cnt int
	for _, v := range intArray {
		if v > 0 { cnt++ }
	}

	fmt.Print(cnt)
}