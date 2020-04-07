package main

import "fmt"

func main(){
	var curNum, b, counter int

	for ;; {
		fmt.Scan(&b)

		if b == 0 {
			break
		}

		if b > curNum  {
			curNum = b
			counter = 1
		} else if b == curNum {
			counter++
		}
	}

	fmt.Println(counter)
}