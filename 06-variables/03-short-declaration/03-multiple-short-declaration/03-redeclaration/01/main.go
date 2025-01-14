package main

import "fmt"

func main() {
	// `safe`'s value is `false`
	var safe bool

	// `safe`'s value is now `true`

	// `speed` is declared and initialized to `50`

	// redeclaration only works when
	//
	// at least one of the variables
	// should be a new variable

	// speed:= 32

	safe, speed, king := true, "asd", "asd"

	fmt.Println(safe, speed, king)

	// EXERCISE
	//
	// Declare the speed variable before
	// the short declaration "again"
	//
	// Observe what happens
}
