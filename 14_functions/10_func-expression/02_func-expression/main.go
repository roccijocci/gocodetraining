 package main

import "fmt"

func main() {

	greeting := func() {
		fmt.Println("Hello world!")
	}

	greeting()
}
// assigning a function to a variable is func expression