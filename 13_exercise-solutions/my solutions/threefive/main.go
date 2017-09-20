package main

import "fmt"

// var a, b, c int

func main() {
	a := 0
	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			a += i
			// fmt.Println(i, "ends")
		} else if i%5 == 0 {
			a += i
			// fmt.Println(i)
		}
	}
	fmt.Println(a)
}
