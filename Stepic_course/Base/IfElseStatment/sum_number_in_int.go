package main

import (
	"fmt"
)

func main() {
	var a int

	fmt.Scan(&a)

	//// example 123456
	i1 := a / 100000 // 1
	i2 := a / 10000 % 10 // 2
	i3 := a / 1000 % 10 // 3
	i4 := a / 100 % 10 // 4
	i5 := a / 10 % 10 // 5
	i6 := a % 10 // 6

	res1 := i1 + i2 + i3
	res2 := i4 + i5 + i6

	if res1 == res2 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}