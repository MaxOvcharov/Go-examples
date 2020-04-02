package main

import "fmt"

func main() {
	var index int

	fmt.Scan(&index)

	var intArray [100]int

	for i := 0; i != index; i++ {
		fmt.Scan(&intArray[i])
	}

	fmt.Println(intArray[index-1])
}
