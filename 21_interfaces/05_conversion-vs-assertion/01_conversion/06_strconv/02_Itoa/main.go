package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 12
	// w := strconv.Itoa(x)
	y := "I have this many: " + strconv.Itoa(x)
	fmt.Println(y)
	// fmt.Printf("%T\n", w)
	// fmt.Println(w)
	// fmt.Println("I have this many: ", strconv.Itoa(x), x)
}
