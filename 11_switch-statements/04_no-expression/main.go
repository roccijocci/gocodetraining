package main

import "fmt"

// func main() {

// 	myFriendsName := "Mar"

// 	switch {
// 	case len(myFriendsName) == 2:
// 		fmt.Println("Wassup my friend with name of length 2")
// 	case myFriendsName == "Tim":
// 		fmt.Println("Wassup Tim")
// 	case myFriendsName == "Jenny":
// 		fmt.Println("Wassup Jenny")
// 	case myFriendsName == "Marcus", myFriendsName == "Medhi":
// 		fmt.Println("Your name is either Marcus or Medhi")
// 	case myFriendsName == "Julian":
// 		fmt.Println("Wassup Julian")
// 	case myFriendsName == "Sushant":
// 		fmt.Println("Wassup Sushant")
// 	default:
// 		fmt.Println("nothing matched; this is the default")
// 	}
// }

/*
  expression not needed
  -- if no expression provided, go checks for the first case that evals to true
  -- makes the switch operate like if/if else/else
  cases can be expressions
*/
func main() {
	var days string
	fmt.Print("Enter a month and determine your horoscope: ")
	fmt.Scan(&days)

	switch {
	case days == "january":
		fmt.Println("U are a jan kid")
	case days == "febuary":
		fmt.Println("valentine baby")
	case days == "march":
		fmt.Println("The Ides of March fall on you")
	case days == "april":
		fmt.Println("Aries born")
	case days == "may":
		fmt.Println("Mars is your abode")
	case days == "june":
		fmt.Println("Peaceful as the dove")
	case days == "july":
		fmt.Println("Fireworks like the fourth of july")
	case days == "august":
		fmt.Println("august babies are the bomb")
	case days == "september":
		fmt.Println("Rule in september")

	default:
		fmt.Println("your case is different")
	}
}
