package main

import "fmt"

type foo int
type bar string

func main() {
	var myAge foo
	myAge = 44
	fmt.Printf("%T %v \n", myAge, myAge)

	var yourAge int
	yourAge = 29
	fmt.Printf("%T %v \n", yourAge, yourAge)

	var yourWord bar
	yourWord = "Long live the queen"
	fmt.Printf("%T %v \n", yourWord, yourWord)

	var myWord string
	myWord = "All hail the king"
	fmt.Printf("%T %v \n", myWord, myWord)
	// this doesn't work:
	// fmt.Println(myAge + yourAge)
	fmt.Println(bar(myWord) + yourWord)
	// this conversion works:
	fmt.Println(int(myAge) + yourAge)
}
