package main
import "fmt"

func main()  {
	array := [5]int{}
	var a int
	for i:=0; i < 5; i++{
		fmt.Scan(&a)
		array[i] = a
	}

	var maxNum int
	for _, i := range array {
		if i > maxNum {
			maxNum = i
		}
	}

	fmt.Println(maxNum)
}