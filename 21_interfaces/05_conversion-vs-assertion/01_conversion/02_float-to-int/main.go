package main

import "fmt"

func main() {
	var x = 12
	var y = 12.1230123
	fmt.Printf("%T\n", y)
	fmt.Println(int(y) + x)
	// conversion: float64 to int
}
