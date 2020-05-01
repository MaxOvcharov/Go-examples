package main

import (
	"fmt"
)

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	var a int
	fmt.Scan(&a)
	var res string
	for ;; {

		if a == 0 {
			break
		}

		if a % 2 == 0 {
			a = a / 2
			res += "0"
		} else {
			a = a / 2
			res += "1"
		}
	}

	fmt.Println(Reverse(res))
}
