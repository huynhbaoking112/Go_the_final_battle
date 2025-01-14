package main

// VARIABLE NAMING RULES

import "fmt"

func main() {
	// CORRECT DECLARATIONS
	var speed int
	var SpeeD int

	// underscore is allowed but not recommended
	var _speed int

	// Unicode letters are allowed
	// var 速度 int

	var king string = "1"

	// keep the compiler happy
	_, _, _, _ = speed, SpeeD, _speed, king

	var son = king

	fmt.Println(king)
	fmt.Println(son)

	king = "a"
	fmt.Println(king)
	fmt.Println(son)

}
