package main

import "fmt"

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	var y int64
	fmt.Println(factorial(4))
	fmt.Print("Enter the number you want to find: ")
	fmt.Scan(&y)
	fmt.Println("The factorial of the number", y, "is", fact(y))
}

func fact(x int64) int64 {
	if x == 0 {
		return 1
	}
	return x * fact(x-1)
}
