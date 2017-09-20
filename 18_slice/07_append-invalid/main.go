package main

import "fmt"

func main() {

	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "buenos dias!"
	// greeting[3] = "suprabadham" but fixed
	greeting = append(greeting, "suprabadham")

	fmt.Println(greeting[3])
	// fmt.Println(len(greeting))
	// fmt.Println(cap(greeting))
}