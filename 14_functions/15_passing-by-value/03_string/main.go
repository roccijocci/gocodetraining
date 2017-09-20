package main

import "fmt"

func main() {
	// just made it basically interactive
	name := "Todd"
	fmt.Print("enter name")
	fmt.Scan(&name)
	// fmt.Println(name) // Todd

	changeMe(name)

	fmt.Println(name) // Todd
}

func changeMe(z string) {
	// fmt.Println(z) // Todd
	fmt.Print("Enter your name")
	fmt.Scan(&z)
	// z = "Rocky"
	// fmt.Println(z) // Rocky
}
