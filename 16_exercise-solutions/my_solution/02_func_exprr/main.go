package main

import "fmt"

func main() {
	test := func(x int) (float32, bool) {
		return float32(x) / 2, x%2 == 0
	}
	w, q := test(13)
	fmt.Println(w, q)
}

// modify
