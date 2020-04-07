package main

import "fmt"

func main(){
	var a, c int = 8, 0
	const b int = 10
	a = a + b
	c = b + a
	fmt.Println(a, c)
}