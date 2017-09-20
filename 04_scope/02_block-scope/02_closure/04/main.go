package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}

func ww() func() int {
	w := 0
	return func() int {
		w++
		return w
	}
}
func test() {
	increament := ww()
	fmt.Println(increament())
	fmt.Println(increament())

}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope
*/
