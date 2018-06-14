package main

import "fmt"
import "math"

const s string  = "Test constant"

func main()  {
	fmt.Println("CONSTANT STRING: ", s)

	const n  = 500
	const d  =3e20 / n

	fmt.Println("CONSTANT INT: ", d)
	fmt.Println("CONSTANT INT64: ", int64(d))

	fmt.Println("MATH - SIN(INT): ", math.Sin(d))
	fmt.Println("MATH - SIN(INT64): ", int64(math.Sin(d)))
}
