package main

import "fmt"

func main() {

	var myGreeting = make(map[string]string)
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)

	//int example of a map
	var testNumbers = make(map[int]int)
	testNumbers[1] = 10000
	testNumbers[2] = 2000

	fmt.Println(testNumbers)
}
