package main

import "fmt"

func main() {
	fmt.Println("\n## FOR LOOP EXAMPLES ##\n" )
	var i int = 0
	for i<3 {
		fmt.Println("FOR LOOP(ver_1): ", i)
		i ++
	}
	fmt.Println("\n###\n" )
	for j := 0; j <= 3; j++ {
		fmt.Println("FOR LOOP(ver_2): ", j)
	}
	fmt.Println("\n###\n" )
	for {
		fmt.Println("FOR LOOP(ver_3): ", "inf loop")
		break
	}
}