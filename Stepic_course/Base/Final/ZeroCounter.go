package main

import "fmt"

func main() {
	var input_num, b int

	fmt.Scan(&input_num)

	cnt := 0
	for i := 1; input_num >= i; i++ {

		fmt.Scan(&b)
		if b == 0 { cnt++ }
	}

	fmt.Println(cnt)
}
