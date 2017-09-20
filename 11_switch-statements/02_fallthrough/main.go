package main

import "fmt"

func main() {
	switch "tim" {
	case "Tim", "tim":
		fmt.Println("Wassup Tim")
	case "Jenny", "jenny":
		fmt.Println("Wassup Jenny")
	case "Marcus", "marcus":
		fmt.Println("Wassup Marcus")
		fallthrough
	case "Medhi", "medhi":
		fmt.Println("Wassup Medhi")
		fallthrough
	case "Julian", "julian":
		fmt.Println("Wassup Julian")
	case "Sushant", "sushant":
		fmt.Println("Wassup Sushant")
	}
}
