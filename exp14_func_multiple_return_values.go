package main

import "fmt"

func plus_or_minus(a, b int) (int, int)  {
	return a + b, a - b
}

func main()  {
	a, b := 5, 6
	p, m := plus_or_minus(a, b)
	fmt.Printf("FUNC_MULTIPLE_RES: (%d, %d)", p, m)
}
