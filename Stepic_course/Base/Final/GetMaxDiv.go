package main

import "fmt"

func main() {
	var a, b int

	fmt.Scan(&a, &b)
	dNum := 7
	if b < dNum || (b - a < dNum && a % dNum != 0 && b % dNum != 0) {
		fmt.Println("NO")
	} else if b % dNum == 0{
		fmt.Println(b / dNum * dNum)
	} else {
		fmt.Println(b - b %dNum)
	}
}
