package main
import "fmt"

func main()  {
	arrayInt := [3]int{}
	var cur, maxNumInx, maxNum int
	for i:=0; i <= 2; i++{
		fmt.Scan(&cur)
		arrayInt[i] = cur

		if cur > maxNum {
			maxNum = cur
			maxNumInx = i
		}
	}

	var a, b int
	if maxNumInx == 0 {
		a = arrayInt[1]
		b = arrayInt[2]
	} else if maxNumInx == 1 {
		a = arrayInt[0]
		b = arrayInt[2]
	} else {
		a = arrayInt[0]
		b = arrayInt[1]
	}

	res := maxNum*maxNum == a*a + b*b
	if res {
		fmt.Println("Прямоугольный")
	} else {
		fmt.Println("Непрямоугольный")
	}
}