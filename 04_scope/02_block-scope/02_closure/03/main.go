package main

import "fmt"

func main() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

func test() {
	y := 0
	increase := func() int {
		y++
		return y
	}
	fmt.Println(increase())
	fmt.Println(increase())
}

func test1() {
	z := 0
	incresed := func() int {
		z++
		return z
	}
	fmt.Println(incresed())
	fmt.Println(incresed())
}

/*
closure helps us limit the scope of variables used by multiple functions
without closure, for two or more funcs to have access to the same variable,
that variable would need to be package scope

anonymous function
a function without a name

func expression
assigning a func to a variable

Todo: delve debugger GOPATH/bin
Install from https://github.com/derekparker/delve
*/
