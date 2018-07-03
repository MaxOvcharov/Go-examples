package main

import (
"fmt"
"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
	//or short way
	i := 42
	fl := float64(i)
	u := uint(f)
	fmt.Println(i, fl, u)
}