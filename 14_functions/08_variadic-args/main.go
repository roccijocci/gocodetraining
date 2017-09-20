package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57}
	info := []string{"sandy", "boy", "meets", "world"}
	// fmt.Println(info)
	// fmt.Printf("the type for data is:  %T \n",data)
	n := average(data...)
	w := stream(info...)
	fmt.Println(n)
	fmt.Println(w)
}

func average(sf ...float64) float64 {
	fmt.Printf("%T The type for sf is : \n", sf)
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
func stream(str ...string) string {
	fmt.Printf("%T the type of str is : \n",str )
	var join string
	for _, x := range str {
		join += x
	}
	return join
}
