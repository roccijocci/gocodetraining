package main

import "fmt"

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

// personal example of callback

func dogs(breed []string, name func(string)) {
	for _, n := range breed {
		name(n)
	}
}

func names(list []string, callback func(string)) {
	for _, w := range list {
		callback(w)
	}
}
func mems(arg []string, mens2 func(string)) {
	for _, m := range arg {
		mens2(m)
	}
}

func main() {
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})

	names([]string{"toyo,adam,hawwa,otto"}, func(w string) {
		fmt.Println(w)
	})

	dogs([]string{`rotweiller,mongrel,pitbull,mastiff`}, func(n string) {
		fmt.Println(n)
	})

	dogs([]string{`pow-pow`, `terrier`, `dalmatian`}, func(n string) {
		fmt.Println(n)
	})

	mems([]string{`poorgi porrgisays poorgidoes`}, func(z string) {
		fmt.Println(z)
	})
}

// callback: passing a func as an argument
