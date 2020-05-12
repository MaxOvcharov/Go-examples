package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a string
	//fmt.Scan(&a)
	a = "1 2 3 5 6"
	aStr := strings.Split(a, " ")
	aInt := make([]int, 0, len(aStr))
	for _, a := range aStr {
		i, _ := strconv.Atoi(a)
		aInt = append(aInt, i)
	}

	fmt.Println(sumInt(aInt...))
}

func sumInt(a ...int) (cnt int, res int) {
	for _, i := range a {
		res += i
		cnt++
	}

	return cnt, res
}