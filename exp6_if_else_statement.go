package main

import "fmt"

func main() {

	if 7%2 == 0 {
		fmt.Println("7 четное число")
	} else {
		fmt.Println("7 нечетное число")
	}


	if 8%4 == 0 {
		fmt.Println("8 делится на 4")
	}


	if num := 9; num < 0 {
		fmt.Println(num, "отрицательное число")
	} else if num < 10 {
		fmt.Println(num, "состоит из 1 цифры")
	} else {
		fmt.Println(num, "состоит из нескольких цифр")
	}
}
