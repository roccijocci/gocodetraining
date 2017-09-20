package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}
type dog struct {
	person
	breed string
	fur   string
}

func main() {
	p1 := person{"James", "Bond", 20}
	p2 := person{"Miss", "Moneypenny", 18}
	// going step ahead and combining struct
	d1 := dog{
		person{
			"suppy",
			"doggie",
			2,
		}, "Mastiff", "short"}
	fmt.Println("the breed is-"+d1.breed, "name of the pet-"+d1.first, "and it's of "+d1.fur+" fur")
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}
