package main

import "fmt"

func greatest(num ...int) int {
	var highest int
	for i, w := range num {
		if w > highest || i == 0 {
			highest = w
		}
	}
	return highest
}

func main() {
	largest := greatest(1, 12, 13, 13, 43, 535, 3647, 758698, 32)
	fmt.Println(largest)
}


func testgreatest(num ...int) int{
	var highest int
	for i, w := range num {
		if w > highest {
			highest = w
		}
	}
	return highest
}