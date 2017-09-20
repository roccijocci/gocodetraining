package main

import "fmt"

func main() {
	for i := 65; i <= 122; i++ {
		fmt.Println(i, " - ", string(i), " - ", i%12)
	}

	test := 65
	tes := string(test)
	fmt.Printf("%T \n", test)
	fmt.Println(tes)

	for index := 0; index < 20; index++ {
		fmt.Printf("%T - ", index)
		fmt.Println(index, "-", string(index))
	}
}
