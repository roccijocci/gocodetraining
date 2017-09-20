package main

import "fmt"

func zero(z *int) {
	*z = 0
}

func main() {
	x := 5
	// fmt.Println(&x) for testing purposes
	zero(&x)
	fmt.Println(x) // x is 0
}
