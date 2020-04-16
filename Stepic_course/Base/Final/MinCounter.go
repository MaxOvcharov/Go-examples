package main

import "fmt"

func main() {
	var input_num, cnt, b int

	fmt.Scan(&input_num)
	if input_num != 0 {
		fmt.Scan(&b)
		cnt = 1
		minNum := b

		for i := 2; input_num >= i; i++ {
			fmt.Scan(&b)
			if minNum > b {
				minNum = b
				cnt = 1
			} else if minNum == b {
				cnt++
			}
		}

		fmt.Println(cnt)
	} else {
		fmt.Println(cnt)
	}
}
