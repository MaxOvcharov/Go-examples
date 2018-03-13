package main

import "fmt"

type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width * r.height
}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func main()  {
	r := rect{width: 10, height: 5}

	fmt.Printf("Area_1: %d\n", r.area())
	fmt.Printf("Perimetr_1: %d\n", r.perim())

	rp := &r
	fmt.Printf("\nArea_2: %d\n", rp.area())
	fmt.Printf("Perimetr_2: %d\n", rp.perim())

}
