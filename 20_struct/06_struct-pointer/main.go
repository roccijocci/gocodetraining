package main

import "fmt"

type person struct {
	name string
	age  int
}
type dog struct {
	person
	breed string
	age   int
	fur   string
}

// tried to see how combining structs playout for pointers
func main() {
	p1 := &person{"James", 20}
	fmt.Println(p1)
	fmt.Printf("%T\n", p1)
	fmt.Println(p1.name)
	fmt.Println(p1.age)

	d1 := &dog{person{
		"canis specie", 12,
	},
		"mastiff", 16, "short and shiny"}
	fmt.Println(d1)
	fmt.Printf("%T \n", d1)
	fmt.Println(d1.breed, d1.age, d1.person.age, d1.fur)
}
