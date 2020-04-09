package main

import "fmt"

func main() {
	var a int

	fmt.Scan(&a)

	firstNum := (a % 10)
	secondNum := (a / 10 % 10)
	thirdNum := (a / 100)

	fmt.Printf("%d%d%d", firstNum, secondNum, thirdNum)
}
