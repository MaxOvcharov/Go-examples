package main

import "fmt"

func main() {
	var a, b int

	fmt.Scan(&a)
	fmt.Scan(&b)

	res := 0

	for i := a; i <= b; i++ {
		res = res + i
	}

	fmt.Println(res)
}