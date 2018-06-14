package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main(){
	i := 1

	fmt.Printf("INTEGER VAL: %d\n", i)

	zeroval(i)
	fmt.Printf("FUNC VAL RESULT: %d\n", i)

	zeroptr(&i)
	fmt.Printf("FUNC POINTER RESULT: %d\n", i)

	fmt.Printf("POINTER ADRESS: %d\n", &i)
}