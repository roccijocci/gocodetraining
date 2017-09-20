package main

import "fmt"

func main() {
	word := "Hello"
	letter := rune(word[0])
	fmt.Println(letter)

	//trying out my examples
	sentence := "This is how we do it"
	sword := rune(sentence[4])
	fmt.Println(sword)
}
