package main

import "fmt"

type square struct {
	side float64
}

type octagon struct {
	side, width float64
}

func (z square) area() float64 {
	return z.side * z.side
}
func (w octagon) area() float64 {
	return w.side * w.side * w.side / w.width
}
func (w octagon) circ() float64 {
	return w.side * w.width
}
func (z square) circ() float64 {
	return z.side
}

type shape interface {
	area() float64
}
type polygon interface {
	circ() float64
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
	fmt.Printf("%T\n", z)
}
func details(q polygon) {
	fmt.Println(q)
	fmt.Println(q.circ())
	fmt.Printf("%T\n", q)
}

func main() {
	s := square{10}
	a := octagon{10, 2}
	// for  area method
	// area 100
	// area 500
	info(s)
	info(a)
	//for circ method
	//circ 20
	//circ 10
	details(a)
	details(s)
	// testing to call the area directly
	// fmt.Println(s.area())
	// fmt.Println(a.circ())
	// fmt.Printf("%T\n", s)
	// fmt.Printf("%T\n", a)
}
