package main

import "fmt"

func reciver(jobs chan int, done chan bool) (){
	for {
		j, more := <-jobs
		if more {
			fmt.Println("received job", j)
		} else {
			fmt.Println("received all jobs")
			done <- true
			return
		}
	}
}

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go reciver(jobs, done)

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	<-done
}