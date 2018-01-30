package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main()  {
	a, b, c := 2, 3, 1
	fmt.Printf("FUNC1: %d + %d = %d\n", a, b, plus(a, b))
	fmt.Printf("FUNC2: %d + %d + %d = %d\n", a, b, c, plusPlus(a, b, c))
}