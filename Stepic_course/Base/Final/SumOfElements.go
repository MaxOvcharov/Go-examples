package main

import "fmt"

func main() {
	var a int

	fmt.Scan(&a)

	res := (a % 10) + (a / 10 % 10) + (a / 100)

	fmt.Println(res)
}
