package main

import "fmt"

type vehicle struct {
	Seats    int
	MaxSpeed int
	Color    string
	Engine   bool    
}

type motocycle struct {
	vehicle
	Wheel int
	Doors int
	// engine bool
}

type car struct {
	vehicle
	Wheels int
	Doors  int
}

type plane struct {
	vehicle
	Jet bool
}

type boat struct {
	vehicle
	Length int
}

func main() {
	prius := car{}
	tacoma := car{}
	bmw528 := car{}
	cars := []car{prius, tacoma, bmw528}

	ducati := motocycle{}
	honda := motocycle{}
	yahama := motocycle{}
	harley := motocycle{}
	bikes := []motocycle{ducati, honda, yahama, harley}

	boeing747 := plane{}
	boeing757 := plane{}
	boeing767 := plane{}
	planes := []plane{boeing747, boeing757, boeing767}

	sanger := boat{}
	nautique := boat{}
	malibu := boat{}
	boats := []boat{sanger, nautique, malibu}

	fmt.Println("CARS")
	for key, value := range cars {
		fmt.Println(key, " - ", value)
	}

	fmt.Println("PLANES")
	for key, value := range planes {
		fmt.Println(key, " - ", value)
	}

	fmt.Println("BOATS")
	for key, value := range boats {
		fmt.Println(key, " - ", value)
	}
	fmt.Println("Bikes")
	for key, value := range bikes {
		fmt.Println(key, " - ", value)
	}
}
