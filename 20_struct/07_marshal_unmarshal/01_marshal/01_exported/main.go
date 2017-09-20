package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First       string
	Last        string
	Age         int
	notExported int
}

func main() {
	p1 := person{"James", "Bond", 20, 007}
	p2 := person{"aldread", "ofsteven", 10, 10}
	bs, _ := json.Marshal(p1)
	ss, _ := json.Marshal(p2)
	fmt.Println(bs)
	fmt.Println(ss)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs))
	fmt.Println(string(ss))
}
