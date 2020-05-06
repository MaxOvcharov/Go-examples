package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Scan(&x, &y, &z)
	fmt.Println(vote(x, y, z))
}

func vote(x int, y int, z int) int {
	var zeroCnt, oneCnt int
	member := [3]int{x, y, z}
	zero, one := 0, 1
	for _, m := range member {
		switch {
		case m == one:
			oneCnt++
		case m == zero:
			zeroCnt++
		}
	}

	if zeroCnt >= oneCnt {
		return zero
	}

	return one
}