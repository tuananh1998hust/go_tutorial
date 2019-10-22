package main

import (
	"fmt"
	"math"
)

// Shape :
type Shape interface {
	area() float64
}

// Circle :
type Circle struct {
	x, y, radius float64
}

// Rectangle :
type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func main() {
	cir := Circle{x: 0, y: 0, radius: 5}
	rec := Rectangle{width: 20, height: 10}

	fmt.Printf("Circle Area %f\n", cir.area())
	fmt.Printf("Reatangle Area %f\n", rec.area())
}
