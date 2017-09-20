package main

import "fmt"

//  switch on types
//  -- normally we switch on value of variable
//  -- go allows you to switch on type of variable

// type contact struct {
// 	greeting string
// 	name     string
// }

// // SwitchOnType works with interfaces
// // we'll learn more about interfaces later
// func SwitchOnType(x interface{}) {
// 	switch x.(type) { // this is an assert; asserting, "x is of this type"
// 	case int:
// 		fmt.Println("int")
// 	case string:
// 		fmt.Println("string")
// 	case contact:
// 		fmt.Println("contact")
// 	default:
// 		fmt.Println("unknown")

// 	}
// }

// func main() {
// 	SwitchOnType(7)
// 	SwitchOnType("McLeod")
// 	var t = contact{"Good to see you,", "Tim"}
// 	SwitchOnType(t)
// SwitchOnType(t.greeting)
// SwitchOnType(t.name)
// }

type dog struct {
	breed string
	age   int
	size  int
	tail  bool
	fur   string
}

// SwitchOnType an exported switch on types
func SwitchOnType(y interface{}) {
	switch y.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case bool:
		fmt.Println("bool")
	case dog:
		fmt.Println("dog")
	default:
		fmt.Println("unknown")
	}
}
func main() {
	SwitchOnType(7)
	SwitchOnType("Name")
	var suppy = dog{"Mastiff", 9, 128, true, "brown"}
	SwitchOnType(suppy)
	SwitchOnType(suppy.fur)
	SwitchOnType(suppy.breed)
	SwitchOnType(suppy.tail)
}
