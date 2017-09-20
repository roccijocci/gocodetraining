package vis

import "fmt"

// PrintVar is exported because it starts with a capital letter
func PrintVar() {
	fmt.Println(MyName)
	fmt.Println(yourName)
}

func Properties() {
	myAccount := "BlockChain"
	myAddress := "zggshsdjknsfu12342hkrdvbher"

	fmt.Println(myAccount)
	fmt.Println(myAddress)
}
