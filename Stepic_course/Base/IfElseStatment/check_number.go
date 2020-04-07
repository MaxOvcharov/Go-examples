package main

import "fmt"

func main() {
	var a int

	fmt.Scan(&a)

	var b, c, d int

	// example 123
	b = a % 100 % 10 // 3
	c = a / 10 % 10 // 2
	d = a / 100 // 1

	if b != c && c != d && b != d {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}