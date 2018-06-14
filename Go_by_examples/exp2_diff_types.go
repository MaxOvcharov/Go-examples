package main

import "fmt"

func main() {
	fmt.Println("STRING: Test " + "string " + "concatenation")

	fmt.Println("INT: " + "1 + 1 = ", 1+1)
	fmt.Println("FLOAT: " + "3.0 / 2.0 = ", 3.0/2.0)

	fmt.Println("BOOL: " + "TRUE and FALSE ->", true && false)
	fmt.Println("BOOL: " + "TRUE or FALSE ->", true || false)
	fmt.Println("BOOL: " + "NOT TRUE ->", !true)

}