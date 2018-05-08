package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Square struct {
	side float64
}

type Shape interface {
	area()
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r Square) area() float64 {
	return r.side * r.side
}

func main() {
	r := Square{side: 10}
	c := Circle{radius: 20}
	fmt.Println("Square Area: ", r.area())
	fmt.Printf("Circle Area: %.2f \n", c.area())
}
