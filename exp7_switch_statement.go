package main

import "fmt"
import "time"

func main() {
	i := 2
	fmt.Println("Write ", i, "as ")
	switch i {
	case 1:
		fmt.Println("one\n")
	case 2:
		fmt.Println("two\n")
	case 3:
		fmt.Println("three\n")
	}

	today := time.Now().Weekday()
	switch today {
	case time.Saturday, time.Sunday:
		fmt.Println("It is the weekend - ", today)
	default:
		fmt.Println("It is the weekday - ", today)
	}

	time := time.Now().Hour()
	switch {
	case time < 12:
		fmt.Println("it's before noon - ", time)
	default:
		fmt.Println("it's after noon - ", time)
	}
}
