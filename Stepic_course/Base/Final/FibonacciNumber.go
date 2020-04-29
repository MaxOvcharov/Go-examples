package main

import (
	"fmt"
)

func main() {
	var a, tmp int
	fmt.Scan(&a)
	priv := 1
	cur := 2
	for i := 3;; i++ {
		tmp = priv
		priv = cur
		cur = cur + tmp
		if priv == a {
			fmt.Printf("%d", i)
			break
		} else if priv > a {
			fmt.Println("-1")
			break
		}
	}
}
