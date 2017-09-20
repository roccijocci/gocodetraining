package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"zero":  "Good morning!",
		"one":   "Bonjour!",
		"two":   "Buenos dias!",
		"three": "Bongiorno!",
	}

	fmt.Println(myGreeting)
	delete(myGreeting, "two")
	fmt.Println(myGreeting)

	myTest := map[int]string{}
	myTest[1] = "mfon"
	myTest[2] = "daddy"

	fmt.Println(myTest)
	delete(myTest, 2)
	fmt.Println(myTest)
}
