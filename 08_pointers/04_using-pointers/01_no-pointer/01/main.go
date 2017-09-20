package main

import "fmt"

func zero(z int) {
	z = 0
}
func start(x, v int) int {
	var s = x + v
	return s
}
func stringop(v, t string) string {
	var str string
	str = v + t
	return str

}
func onefunc(a string) string {
	return a + " amen"
}
func numberofYears(a string) string {
	return a + ` years`
}

// had to practice my functions ahead
func main() {
	x := 5
	w := 3
	y := 4
	a := `amen `
	b := `sister`
	c := `we are love`
	var q int
	fmt.Print(`how old are u ?`)
	fmt.Scanln(&q)
	fmt.Println(`i am `, q, ` years old`)

	fmt.Println(onefunc(c))
	fmt.Println(stringop(a, b))
	//start(w, y)
	zero(x)
	fmt.Println(x) // x is still 5
	fmt.Println(start(w, y))
}
