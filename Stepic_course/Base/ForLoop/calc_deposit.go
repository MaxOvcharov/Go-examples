package main

import "fmt"

func main() {
	var x, p, y int

	fmt.Scan(&x, &p, &y)

	count := 0
	for i := 1; x <= y; i++ {
		x = x + (int)(x * p/100)
		count++
	}

	fmt.Println(count)
}
