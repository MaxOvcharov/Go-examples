package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	next := n + 1
	priv := n - 1

	if n > 0 {
		fmt.Print("Positive number. The next number for the number ", n, " is ", next, ".\n")
		fmt.Print("Positive number. The previous number for the number ", n, " is ", priv, ".")
	} else if n < 0{
		fmt.Print("Negative number. The next number for the number ", n, " is ", next, ".\n")
		fmt.Print("Negative number. The previous number for the number ", n, " is ", priv, ".")
	}
}
