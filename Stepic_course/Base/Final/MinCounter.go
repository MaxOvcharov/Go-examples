package main

import (
	"fmt"
)

func main() {
	var a, res int

	fmt.Scan(&res)
	if res > 20 {
		a = res % 10
	} else {
		a = res
	}

	var kw string
	switch {
	case a == 1:
		kw = "korova"
	case a > 1 && a < 5:
		kw = "korovy"
	case a == 0 || a >= 5 && a <= 20:
		kw = "korov"
	}

	fmt.Printf("%d %s", res, kw)
}
