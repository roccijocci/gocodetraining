package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []string{"Zeno", "John", "Al", "Jenny"}
	fmt.Println(s)
	sort.StringSlice(s).Sort()
	fmt.Println(s)
	// learning to access methods of the string slice
	a := sort.StringSlice(s).Len()
	w := sort.StringSlice(s).Less(0, 1)
	q := sort.StringSlice(s).Search("Al")
	sort.StringSlice(s).Swap(0, 1)
	// sort.Sort(sort.StringSlice(s))
	fmt.Println(a)
	fmt.Println(w)
	fmt.Println("the string you requested is in the ", q, " index")
	fmt.Println(s)
}

// https://golang.org/pkg/sort/#Sort
