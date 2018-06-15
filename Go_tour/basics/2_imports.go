package main

// Bad way of using imports

//import "fmt"
//import "math"

// God way of using imports
import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.", math.Sqrt(7))
}
