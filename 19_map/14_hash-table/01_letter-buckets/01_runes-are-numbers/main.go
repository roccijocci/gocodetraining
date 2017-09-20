package main

import "fmt"

func main() {
	letter := 'A'
	fmt.Println(letter)
	fmt.Printf("%T \n", letter)

	var let byte
	let = 'C'
	fmt.Printf("%T \n", let)
	fmt.Println(let)
}
