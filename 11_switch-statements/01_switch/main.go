package main

import "fmt"

// func main() {
// 	switch "Jenny" {
// 	case "Daniel":
// 		fmt.Println("Wassup Daniel")
// 	case "Medhi":
// 		fmt.Println("Wassup Medhi")
// 	case "Jenny":
// 		fmt.Println("Wassup Jenny")
// 	default:
// 		fmt.Println("Have you no friends?")
// 	}
// }

/*
  no default fallthrough
  fallthrough is optional
  -- you can specify fallthrough by explicitly stating it
  -- break isn't needed like in other languages
*/

func main() {
	switch "Boy" {
	case "boy":
		fmt.Println("i am a boy")
	case "girl":
		fmt.Println("i am a girl")
	case "woman":
		fmt.Println("woman of virtue")
	default:
		fmt.Println("have you no virtue")
	}
}
