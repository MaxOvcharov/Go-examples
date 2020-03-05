package main

import "fmt"

func main() {
	var a int

	fmt.Scan(&a)

	var cur, priv int

	priv = a

	for ;; {
		cur = priv / 10
		if cur > 0 {
			priv = cur
		} else {
			break
		}
	}

	fmt.Println(priv)
}