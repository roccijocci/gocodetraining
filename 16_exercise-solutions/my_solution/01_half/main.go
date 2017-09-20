package main

import (
	"fmt"
)

//contains i and ii together
func pers(x int) (int, bool) {
	return x / 2, x%2 == 0
}
func assignment(y int) (float32, bool) {
	return float32(y) / 2, y%2 == 0
}

func main() {
	b, r := pers(4)
	fmt.Println(b, r)
	w, m := assignment(19)
	fmt.Println(w, m)
}
