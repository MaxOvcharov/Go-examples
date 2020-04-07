package main

import (
	"fmt"
	"strings"
)

func main() {
	var index int

	fmt.Scan(&index)

	var strArray [100]string

	var result []string
	for i := 0; i != index; i++ {
		fmt.Scan(&strArray[i])
	}

	for i, v := range strArray {
		if i % 2 == 0 {
			result = append(result, v)
		}
	}

	fmt.Println(strings.Join(result, " "))
}