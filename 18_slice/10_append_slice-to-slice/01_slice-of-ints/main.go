package main

import "fmt"

func main() {

	// mySlice := []int{1, 2, 3, 4, 5}
	// myOtherSlice := []int{6, 7, 8, 9}

	// mySlice = append(mySlice, myOtherSlice...)

	// fmt.Println(mySlice)

	newSlice := []int{0, 0, 1, 1, 2, 3}
	otherSlice := []int{61, 2, 3, 4}

	// otherSlice = append(otherSlice, newSlice...)
	newSlice = append(newSlice, otherSlice...)

	fmt.Println(newSlice)
}
