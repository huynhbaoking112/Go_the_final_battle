package main

import "fmt"

var declareMeAgain = 10

func nested() {

	//declares the same variable
	// they both can exist together
	// this one only belongs to this scope
	// package's variable is still intact

	var declareMeAgain = 5
	fmt.Println("inside Nested: ", declareMeAgain)
}

func main() {

	fmt.Println("Inside main:", declareMeAgain)

	nested()
	// package-level declareMeAgain isn't effected
	// from the change inside the nested func

	fmt.Println("inside main:", declareMeAgain)

}
