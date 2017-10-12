package main

import "fmt"

type animal struct {
	sound string
	fur   bool
	name  string
}

type dog struct {
	animal
	friendly bool
}

type cat struct {
	animal
	annoying bool
}

func specs(a interface{}) {
	fmt.Println(a)
	fmt.Printf("%T  \n", a)
}

func main() {
	fido := dog{animal{"woof", false, "offsteven"}, true}
	fifi := cat{animal{"meow", true, "offgwen"}, true}
	specs(fido)
	specs(fifi)
}
