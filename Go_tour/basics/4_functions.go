package main

import "fmt"

func add(x, y int) int {
	return (x + y) / 2
}

func main() {
	fmt.Println(add(42, 13))
}