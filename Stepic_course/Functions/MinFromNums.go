package main

import "fmt"

func main() {
	var a int

	fmt.Scan(&a)
	fmt.Println(fibonacci(a))
}

func fibonacci(n int) int {
	priv, cur := 1, 1
	for i := 1; i < n; i++ {
		priv, cur = cur, priv + cur
	}

	return priv
}