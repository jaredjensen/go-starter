package main

import (
	"fmt"
	"math"
)

// Shape is a shape
type Shape interface {
	area() float64
}

// Circle is a shape
type Circle struct {
	x, y, radius float64
}

// Rectangle is a shape
type Rectangle struct {
	height, width float64
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 10}
	fmt.Println(circle, getArea(circle))

	rectangle := Rectangle{height: 10, width: 20}
	fmt.Println(rectangle, getArea(rectangle))

}
