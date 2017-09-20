package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

// another shape
type circle struct {
	radius float64
}
type rect struct {
	length, breath float64
}

type shape interface {
	area() float64
}

func (r rect) area() float64 {
	return r.length * r.breath
}
func (s square) area() float64 {
	return s.side * s.side
}

// which implements the shape interface
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := square{10}
	c := circle{5}
	r := rect{2, 4}
	info(s)
	info(c)
	info(r)
}
