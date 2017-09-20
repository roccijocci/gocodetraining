package main

import "fmt"

func main() {
	n := hashBucket("Ao", 22)
	fmt.Println(n)

	w := testhash("Bug", 10)
	fmt.Println(w)
}

func hashBucket(word string, buckets int) int {
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}
func testhash(words string, buck int) int {
	let := int(words[0])
	butt := let % buck
	return butt
}
