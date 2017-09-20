package main

import (
	"fmt"
)

type person struct {
	Name string
	Age  int
}

// to make it clear in my eyes
type dog struct {
	name  string
	breed string
	age   int
}

type animal struct {
	dog
	endangered bool
	habitat    string
}

type doubleZero struct {
	person
	LicenseToKill bool
}

func (d dog) Speech() {
	fmt.Println("I am a dog")
}
func (a animal) Speech() {
	fmt.Println("We are all animals")
}
func (p person) Greeting() {
	fmt.Println("I'm just a regular person.")
}

func (dz doubleZero) Greeting() {
	fmt.Println("Miss Moneypenny, so good to see you.")
}

func main() {
	p1 := person{
		Name: "Ian Flemming",
		Age:  44,
	}

	p2 := doubleZero{
		person: person{
			Name: "James Bond",
			Age:  23,
		},
		LicenseToKill: true,
	}

	d1 := dog{
		name:  "Duke",
		breed: "caucasian",
		age:   2,
	}

	a1 := animal{
		dog: dog{
			name:  "deku",
			breed: "snow monkey",
			age:   4,
		},
		endangered: true,
		habitat:    "mountains of japan",
	}
	p1.Greeting()
	p2.Greeting()
	p2.person.Greeting()
	// testing how it works out
	d1.Speech()
	a1.Speech()
	a1.dog.Speech()
}
