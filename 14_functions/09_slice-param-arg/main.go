package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57}
	list := []string{`i`,`z`,`z`,`l`,`e`}
	n := average(data)
	w := flist(list)
	fmt.Println(n)
	fmt.Println(w)
}

func average(sf []float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
func flist(fl []string) string {
	var con string
	for _, s := range fl{
		con += s
	}
	return con
}
