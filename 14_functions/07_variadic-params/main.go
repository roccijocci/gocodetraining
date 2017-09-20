package main

import "fmt"

func main() {
	n := average(43, 56, 87, 12, 45, 57)
	w := sum(43, 56, 87, 12, 45, 57)
	s := stream(`H`, `E`, `L`, `L`, `O`)
	fmt.Println(s)
	fmt.Printf("%T \n",s )
	
	fmt.Println(w)
	fmt.Println(n)
	fmt.Println(s)
}

func average(sf ...float64) float64 {
	// fmt.Println(sf)
	// fmt.Printf("%T \n", sf)
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

func sum(tt ...int32) int32 {
	// fmt.Println(tt)
	// fmt.Printf("%T \n", tt)
	var summation int32
	for _, e := range tt {
		summation += e
	}
	return summation
}

func stream(st ...string) string {
	fmt.Println(st)
	fmt.Printf("%T \n",st)
	var envtest string
	for _, w := range st {
		envtest += w
	}
	return envtest
}
