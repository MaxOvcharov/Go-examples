package main

import "fmt"

func main()  {

	s1 := []int{1, 2, 3, 4, 5}
	sum := 0
	fmt.Println("SLICE: ", s1)
	for _, num := range s1 {
		sum += num
	}

	fmt.Println("Iterate RANGE nums from SLICE - sum: ", sum )

	for i, num := range s1 {
		if num == 4{
			fmt.Println("Iterate RANGE nums from SLICE - index of num=4: ", i)
		}
	}

	m1 := map[string]int{"k1": 1, "k2": 2, "k3": 3, "k4": 4, "k5": 5}
	fmt.Println("MAP: ", m1)

	for key, value := range m1 {
		fmt.Printf("KEY: %s VALUE: %d\n", key, value)
	}

	str1 := "test"
	fmt.Println("STRING: ", str1)
	for i, c := range str1 {
		fmt.Printf("INDEX: %d CHAR: %d\n", i, c)
	}
}
