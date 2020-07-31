package main

import (
	"fmt"
	"math"
)

// Define Interface
type Shape interface {
	area() float64
}

type Circle struct {
	x      float64
	y      float64
	radius float64
}

type Rectangle struct {
	length, width float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.length * r.width
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	fmt.Println("Hello World")
	circle1 := Circle{x: 5, y: 5, radius: 10}
	rectangle1 := Rectangle{length: 3, width: 4}

	fmt.Printf("Circle area : %f\n", getArea(circle1))
	fmt.Printf("Rectangle area : %f\n", getArea(rectangle1))
}
