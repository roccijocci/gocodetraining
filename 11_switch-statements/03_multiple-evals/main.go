package main

import "fmt"

// func main() {
// 	switch "Jenny" {
// 	case "Tim", "Jenny":
// 		fmt.Println("Wassup Tim, or, err, Jenny")
// 	case "Marcus", "Medhi":
// 		fmt.Println("Both of your names start with M")
// 	case "Julian", "Sushant":
// 		fmt.Println("Wassup Julian / Sushant")
// 	}
// }

func main() {
	var name string
	fmt.Print("What is ur name:")
	fmt.Scan(&name)
	switch name {
	case name:
		fmt.Println("my name is ", name)
	default:
		fmt.Println("guess u missed this")
	}
}
