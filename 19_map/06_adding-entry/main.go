package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
	}
	// elements in the maps are randomly placed
	myGreeting["Harleen"] = "Howdy"
	myGreeting["Joker"] = "Harley Quinn"
	myGreeting["aaaa"] = "bbbb"

	//tested out the learn func
	fmt.Println("The length of this map is", len(myGreeting))
	fmt.Println(myGreeting)
}
