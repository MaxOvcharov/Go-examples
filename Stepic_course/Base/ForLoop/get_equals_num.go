package main

import (
	"fmt"
	"strings"
)

func main(){
	var first, second string

	fmt.Scan(&first)
	fmt.Scan(&second)
	var result []string

	for _, i := range first {
		for _, j := range second {
			if i == j {
				result = append(result, string(i))
			}
		}
	}

	fmt.Println(strings.Join(result, " "))
}