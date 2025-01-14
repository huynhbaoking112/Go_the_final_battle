package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Multiple #2
//
//  1. Declare and initialize two string variables
//     using multiple variable declaration
//
//  2. Use the type once while declaring the variables
//
//  3. The first variable's name should be firstName
//  4. The second variable's name should be lastName
//
//  5. Print them all
//
// EXPECTED OUTPUT
//  "" ""
// ---------------------------------------------------------

func main() {
	var firstName, lastName string = "", ""

	// fmt.Println(firstName, lastName)
	fmt.Printf("%q %q", firstName, lastName)
}
