package main

import "fmt"

func main() {

	age := 44
	fmt.Println(&age) // 0x82023c080

	// fmt.Println("test answer", age) //24|| 44
	changeMe(&age)

	fmt.Println(&age) //0x82023c080
	fmt.Println(age)  //24

	name := 12
	fmt.Println(`prints main`, name)
	fmt.Println("this is a test", &name) // reveals the memory location of 12
	explain(&name)
	fmt.Println(`testing`, name)
}

func changeMe(z *int) {
	fmt.Println(z)  // 0x82023c080
	fmt.Println(*z) // 44
	*z = 24
	fmt.Println(z)  // 0x82023c080
	fmt.Println(*z) // 24
}

func explain(w *int) {
	// fmt.Println(" explain testing ", w)
	// fmt.Println(`another test`, *w)
	*w = 35
	// fmt.Println(`func test`, w)
	// fmt.Println(`func1 test`, *w)

}
