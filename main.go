package main

import (
	"fmt"
)

const PI = 3.14

type shape interface {
	area() float64
	perimeter() float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return PI * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * PI * c.radius
}

type rectangle struct {
	length  float64
	breadth float64
}

func (r rectangle) area() float64 {
	return r.length * r.breadth
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.length + r.breadth)
}
func shapeInfo(s shape) {
	fmt.Printf("Area: %.2f\n", s.area())
	fmt.Printf("Perimeter:%.2f\n", s.perimeter())
}

func main() {
	C := circle{radius: 3.25}
	P := rectangle{length: 5, breadth: 6}
	shapeInfo(C)
	shapeInfo(P)
	fmt.Println("Let start Goroutine")
	Goroutine()
	fmt.Println("All goroutine finished")

}
