package main

import (
	"fmt"
	"sort"
)

type people []string

//type people []string
// studyGroup := people{"Zeno", "John", "Al", "Jenny"}

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
	studyGroup := people{"Zeno", "yvonne", "Gabriel", "john", "Albert", "Jenny"}
	fmt.Println(studyGroup)
	sort.Sort(studyGroup)
	fmt.Println(studyGroup)
}
