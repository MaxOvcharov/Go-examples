package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	oneHour := 3600
	oneMinute := 60

	hours := a / oneHour
	minutes := (a % oneHour) / oneMinute

	fmt.Printf("It is %d hours %d minutes.", hours, minutes)
}
