package main

import "fmt"

const metersToYards float64 = 1.09361
const centi2meters float64 = 0.01
const meters2kilometers = 0.001

// tested a couple of measurement. and included an external function.

func main() {
	var meters float64
	var centi float64
	// hacksaw ridge inspired
	var doss float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, " meters is ", yards, " yards.")
	fmt.Print("Convert Centimeters to meters \n")
	fmt.Print(`Enter Centimeter: `)
	fmt.Scan(&centi)
	meter := centi * centi2meters
	fmt.Println(centi, `centimeters is`, meter, "meters")
	fmt.Print("conversion from meters to kilometers: \n ")
	kilometer(doss)
}

func kilometer(distance float64) {
	fmt.Print("enter distance in meters: ")
	fmt.Scan(&distance)
	calc := distance * meters2kilometers
	fmt.Println(calc, `km`)

}
