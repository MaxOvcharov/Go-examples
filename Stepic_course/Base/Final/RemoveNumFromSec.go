package main

import (
	"fmt"
)

func main() {
	var str, a string
	fmt.Scan(&str, &a)

	for _, c := range(str) {
		cur := string(c)
		if cur != a {
			fmt.Printf("%s", cur)
		}
	}
}
