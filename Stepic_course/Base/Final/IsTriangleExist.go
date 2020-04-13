package main
import "fmt"

func main()  {
	var a, b, c int
	fmt.Scan(&a, &b, &c)

	res1 := a + b > c
	res2 := c + b > a
	res3 := c + a > b


	if res1 && res2 && res3 {
		fmt.Println("Существует")
	} else {
		fmt.Println("Не существует")
	}
}