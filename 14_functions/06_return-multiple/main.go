package main

import "fmt"

func main() {
	fmt.Println(greet("Jane ", "Doe "))
	fmt.Println(school("bank ", "school "))
}

func greet(fname, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}
func school(primary, secondary string) (string, string) {
	return fmt.Sprint(primary, secondary), fmt.Sprint(secondary, primary)
}
