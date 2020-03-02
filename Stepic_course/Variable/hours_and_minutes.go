package main

import "fmt"

func main() {
	var a int

	fmt.Scan(&a)

	oneHour := 360 / 12
	oneMinute := 60 / oneHour

	hours := a / oneHour
	minutes := a % oneHour * oneMinute

	fmt.Println("It is", hours, "hours", minutes, "minutes.")
}
