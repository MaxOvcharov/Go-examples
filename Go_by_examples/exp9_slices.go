package main


import "fmt"

func main()  {
	sl1 := make([]string, 3)
	fmt.Println("EMPTY SLICE: ", sl1, ", EMPTY SLICE LEN: ", len(sl1))

	sl1[0] = "a"
	sl1[1] = "b"
	sl1[2] = "c"
	fmt.Println("SET VALUES: ", sl1, ", GET VALUE: ", sl1[1])

	fmt.Println("SLICE LEN: ", len(sl1))

	sl1 = append(sl1, "d")
	sl1 = append(sl1, "e", "f")
	fmt.Println("APPEND VALUES INTO SLICE:", sl1, ", SLICE LEN: ", len(sl1))

	sl2 := make([]string, len(sl1))
	copy(sl2, sl1)
	fmt.Println("COPY OLD SLICE: ", sl2, ", SLICE LEN: ", len(sl2))

	fmt.Println("SLICE[:5]: ", sl2[:5], ", SLICE[2:]: ", sl2[2:], ", SLICE[2:5]: ", sl2[2:5])

	sl3 := []string{"g", "h", "i"}
	fmt.Println("DECLARED SLICE: ", sl3, ", SLICE LEN: ", len(sl3))

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D SLICE: ", twoD)
}