package main

import "fmt"

func main()  {
	ch_msg := make(chan string)

	go func() {ch_msg <- "ping"} ()

	msg_from_ch := <- ch_msg
	fmt.Printf("Mesage from channel: %s", msg_from_ch)
}
