package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string
	var res int
	fmt.Scan(&a)
	for ;; {
		if res == 0 || res > 9 {
			tmpRes := 0
			for _, char := range a {
				tmpRes += int(char) - 48
			}
			a = strconv.Itoa(tmpRes)
			res = tmpRes
		} else {
			break
		}
	}

	fmt.Println(res)
}
