package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x = "12"
	fmt.Printf("%T\n", x)
	var y = 6
	z, _ := strconv.Atoi(x)
	fmt.Println(y + z)
}
