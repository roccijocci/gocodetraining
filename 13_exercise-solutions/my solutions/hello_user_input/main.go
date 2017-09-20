package main

import "fmt"

var name string

func main() {
	fmt.Print("Enter your Name: ")
	fmt.Scan(&name)
	fmt.Println("Hello ", name)
}
