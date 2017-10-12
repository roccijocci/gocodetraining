package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type rect struct {
	length, breadth float64
}

type shape interface {
	area() float64
	recarea() float64
}

// value receiver example
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) recarea() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rect) recarea() float64 {
	return r.length * r.breadth
}
func (r rect) area() float64 {
	return r.length * r.breadth
}

func info(s shape, a shape) {
	fmt.Println("area of circle ", s.area())
	fmt.Println("area of rectangle ", a.recarea())
}
func main() {
	c := circle{5}
	r := rect{12, 5}
	info(c, r)
}
