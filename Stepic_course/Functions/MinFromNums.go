package main

import "fmt"

func main() {
	fmt.Println(minimumFromFour())
}

func minimumFromFour() int {
	var a, minNum int

	fmt.Scan(&a)
	minNum = a

	for i := 1; i < 4; i++ {
		fmt.Scan(&a)
		if a < minNum {
			minNum = a
		}
	}

	return minNum
}