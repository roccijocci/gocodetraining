package main

import (
	"fmt"
	"sort"
)

type people []string

//type people []string
// studyGroup := people{"Zeno", "John", "Al", "Jenny"}
//s := []string{"Zeno", "John", "Al", "Jenny"}
// n := []int{7, 4, 8, 2, 9, 19, 12, 32,3}

func (p people) Len() int {
	return len(p)
}
func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}
func main() {
	//noticed it treats Uppercase and lowercase differently
	studyGroup := people{"Zeno", "Yvonne", "Gabriel", "John", "Albert", "Jenny"}
	s := []string{"Zeno", "John", "Al", "Jenny"}
	n := []int{7, 4, 8, 2, 9, 19, 12, 32, 3}
	fmt.Println(n)
	fmt.Println(s)
	fmt.Println(studyGroup)
	sort.Ints(n)
	sort.Strings(s)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
	fmt.Println("The sorted Names are", s)
	fmt.Println("The sorted integers are", n)
}
