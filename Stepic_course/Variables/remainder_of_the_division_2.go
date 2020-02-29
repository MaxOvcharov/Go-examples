package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Scan(&a)

	b = a % 100 / 10

	fmt.Println(b)
}
