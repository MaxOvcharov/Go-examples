package main

import "fmt"

func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
    nextInt := intSeq()
    fmt.Println("STEP_1: ", nextInt())
    fmt.Println("STEP_1: ", nextInt())
    fmt.Println("STEP_1: ", nextInt())
    newInts := intSeq()
    fmt.Println("STEP_2: ", newInts())
}