package main

import "fmt"

// func main() {
// 	for i := 50; i <= 140; i++ {
// 		fmt.Printf("%v - %v - %v \n", i, string(i), []byte(string(i)))
// 	}
// }
func main() {
	for i := 50; i <= 100; i++ {
		fmt.Println(i, string(i), []byte(string(i)))
	}
}
