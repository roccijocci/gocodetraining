package main

import "fmt"

func main() {

	// mySlice := []string{"Monday", "Tuesday"}
	myOtherSlice := []string{"Wednesday", "Thursday", "Friday", "saturday", "dimarche"}

	// testSlice := []string{"do", "you", "love", "me"}
	// // testSlice1 := []string{"not", "i"}

	// testSlice = append(testSlice[:2], testSlice[4:]...)
	// fmt.Println(testSlice)

	// mySlice = append(mySlice, myOtherSlice...)
	// fmt.Println(mySlice)

	// mySlice = append(mySlice[:2], mySlice[3:]...)
	// fmt.Println(mySlice)

	myOtherSlice = append(myOtherSlice[:2], myOtherSlice[3:]...)
	fmt.Println(myOtherSlice)

}
