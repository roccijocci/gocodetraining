package main

import "fmt"

func main() {

	name := "Todd"
	fmt.Println("Enter a name") // 0x82023c080
	fmt.Scan(&name)

	changeMe(&name)

	fmt.Println(&name) //0x82023c080
	fmt.Println(name)  //Rocky
}

func changeMe(z *string) {
	fmt.Println(z)  // 0x82023c080
	fmt.Println(*z) // Todd
	// here it changes when inputs
	fmt.Println("enter a change of name")
	fmt.Scan(&*z)
	// *z = "Rocky"
	fmt.Println(z)  // 0x82023c080
	fmt.Println(*z) // Rocky
}
