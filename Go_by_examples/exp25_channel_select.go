package main

import (
	"fmt"
	"time"
)

func f1(c1 chan string) {
	time.Sleep(time.Second * 1)
	c1 <- "FUNC1"
}

func f2(c2 chan string) {
	time.Sleep(time.Second * 2)
	c2 <- "FUNC2"
}

func main () {
	c1 := make(chan string)
	c2 := make(chan string)

	go f1(c1)
	go f2(c2)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Printf("GET(%d): %s\n", i, msg1)
		case msg2 := <-c2:
			fmt.Printf("GET(%d): %s\n", i, msg2)
		}
	}
}