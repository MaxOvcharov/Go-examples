package main

import "fmt"

func main()  {
	var s1 string = "Test string"
	fmt.Println("STRING: ", s1)

	var i1, i2 int = 1, 3
	fmt.Println("INT: ", i1, i2)

	var b1 bool = true
	fmt.Println("BOOL: ", b1)

	var i3 int
	fmt.Println("DEFAULT INT: ", i3)

	s2 := "Test string"
	fmt.Println("SHORT DECLARATION VAR: ", s2)
}
