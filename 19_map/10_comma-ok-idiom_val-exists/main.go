package main

import "fmt"

var elements = map[string]string{
	"H":  "Hydrogen",
	"He": "Helium",
	"Li": "Lithum",
	"Be": "Beryllium",
	"B":  "Boron",
	"C":  "Carbon",
	"N":  "Nitrogen",
	"Fl": "Florine",
	"Ne": "Neon",
}

func main() {

	// myGreeting := map[int]string{
	// 	0: "Good morning!",
	// 	1: "Bonjour!",
	// 	2: "Buenos dias!",
	// 	3: "Bongiorno!",
	// }

	// fmt.Println(myGreeting)

	// // delete(myGreeting, 2)

	// if val, exists := myGreeting[2]; exists {
	// 	fmt.Println("That value exists.")
	// 	fmt.Println("val: ", val)
	// 	fmt.Println("exists: ", exists)
	// } else {
	// 	fmt.Println("That value doesn't exist.")
	// 	fmt.Println("val: ", val)
	// 	fmt.Println("exists: ", exists)
	// }

	// fmt.Println(myGreeting)

	// myTest := make(map[string]string)

	// myTest["igbo"] = "nwanne"
	// myTest["hausa"] = "fatima"
	// myTest["yoruba"] = "ajoke"

	// fmt.Println(myTest)
	// // rem: with this if expression the key is tested
	// // as a bool to check if it's true or false
	// if value, key := myTest["yoruba"]; key {
	// 	fmt.Println("the value exists")
	// 	fmt.Println("val", value)
	// 	fmt.Println("exist", key)
	// } else {
	// 	fmt.Println("the value does not exists")
	// 	fmt.Println("value:", value)
	// 	fmt.Println("exist", key)
	// }
	// fmt.Println(myTest)

	//practiced an example in a book used as a
	//reference guide
	elements := map[string]map[string]string{

		"H": map[string]string{
			"name":  "Hydrogen",
			"state": "gas",
		},
		"He": map[string]string{
			"Name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithum",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"Fl": map[string]string{
			"name":  "Florine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	fmt.Println(elements)
	if key, value := elements["O"]; value {
		fmt.Println(key["name"], key["state"], "does exist")
	} else {
		fmt.Println(key, "does not exist")
	}
	// trying out the delete
	// delete(elements, "O")
	// updating
	// elements["H"] = map[string]string{
	// 	"name":  "mboro",
	// 	"state": "liquid",
	// }

	fmt.Println(elements)
	for key, value := range elements {
		fmt.Println(key, "-", value)
	}
}
