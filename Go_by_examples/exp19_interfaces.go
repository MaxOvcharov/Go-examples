package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}


type circle struct {
	radius float64
}

type rectangle struct {
	width, height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) perimeter() float64 {
	return 2 * r.width + 2 * r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measures(g geometry)  {
	fmt.Printf("Interface data: %f\n", g)
	fmt.Printf("AREA: %f\n", g.area())
	fmt.Printf("PERIMETER: %f\n", g.perimeter())
}

func main()  {
	r := rectangle{1, 3}
	c := circle{5}

	measures(r)
	measures(c)

}