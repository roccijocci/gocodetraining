package main

import "fmt"

var smallnumber, largenumber, remainder int

func main() {
	fmt.Printf("Enter a small number: ")
	fmt.Scanln(&smallnumber)
	fmt.Printf("Enter a large number: ")
	fmt.Scanln(&largenumber)

	remainder = largenumber % smallnumber

	fmt.Println("The remainder of ", largenumber, "divided by", smallnumber, "is", remainder)
}
