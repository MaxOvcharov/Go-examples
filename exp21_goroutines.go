package main

import "fmt"

func f(from string){
	for i := 0; i < 3; i++ {
		fmt.Printf("%s: %d\n", from, i)
	}
}

func main(){
	fmt.Println("START...\n")
	f("sync func")
	go f("async func")

	go func(msg string){
		fmt.Println(msg)
	}("Anonymus async func\n")

	fmt.Scanf("PRESS ENTER:")

	fmt.Println("END...\n")
}

