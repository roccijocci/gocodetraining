package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type dog struct {
	name  string
	breed string
	age   int
}

func (p person) fullName() string {
	return p.first + p.last
}

func (d dog) detail() string {
	return d.name + d.breed
}

func main() {
	p1 := person{"James", "Bond", 20}
	p2 := person{"Miss", "Moneypenny", 18}
	fmt.Println(p1.fullName())
	fmt.Println(p2.fullName())
	d1 := dog{"tutu", "mongrel", 2}
	d2 := dog{"duchess", "mastiff", 3}
	fmt.Println(d1.detail())
	fmt.Println(d2.detail())
}
