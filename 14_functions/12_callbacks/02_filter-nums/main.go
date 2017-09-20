package main

import "fmt"

func filter(numbers []int, callback func(int) bool) []int {
	var xs []int
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	return xs
	//fmt.Printf("%T \n", callback)
}

func test(letters []string, callback func(string) bool) []string {
	var let []string
	for _, l := range letters {
		if callback(l) {
			let = append(let, l)
		}
	}
	return let
}

func main() {
	xs := filter([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})
	fmt.Println(xs) // [2 3 4]

	let := test([]string{`H`, `E`, `L`, `L`, `O`}, func(x string) bool {
		return true
	})
	fmt.Println(let)
}
