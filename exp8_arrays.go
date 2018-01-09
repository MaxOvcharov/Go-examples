package main

import "fmt"

func main()  {
	var a1 [5]int
	fmt.Println("Empty array: ", a1)
	i := 100
	a1[0] = i
	fmt.Println("SET: ", i, ", GET: ", a1[0], ", ARRAY: ", a1)

	fmt.Println("LEN ARRAY: ", len(a1))

	a2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("DECLAR NEW ARRAY: ", a2)

	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("EMBEDDED ARRAY: ", twoD)
}
