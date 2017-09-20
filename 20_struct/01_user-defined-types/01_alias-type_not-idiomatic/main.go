package main

import "fmt"

type foo int
type bar string

func main() {
	var myAge foo
	myAge = 44
	fmt.Printf("%T %v \n", myAge, myAge)

	var myName bar
	myName = "rotciv goes to school"
	fmt.Printf("%T \n", myName)
	fmt.Println(myName)
}
