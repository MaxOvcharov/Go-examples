package main

import "fmt"

func main(){
	var a, b, result int
	fmt.Scan(&a)

	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		if b >= 10 && b <= 99 && b % 8 == 0 {
			result = result + b
		}
	}

	fmt.Println(result)
}