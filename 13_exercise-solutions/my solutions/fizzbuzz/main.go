package main

import "fmt"

func main() {
	for index := 0; index <= 100; index++ {
		if index%15 == 0 {
			fmt.Println("fizzbuzz", index)
		} else if index%3 == 0 {
			fmt.Println("fizz", index)
		} else if index%5 == 0 {
			fmt.Println("Buzz", index)
		} else {
			fmt.Println(index)
		}
	}
}
