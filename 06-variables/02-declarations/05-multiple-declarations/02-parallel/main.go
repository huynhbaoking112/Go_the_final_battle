package main

import "fmt"

func main() {
	// this is equal to:
	//
	//   var (
	//     speed int
	//     velocity int
	//   )
	//
	// or:
	//
	//   var speed int
	//   var velocity int
	//
	// var speed, velocity int

	// fmt.Println(speed, velocity)

	var (
		king, bin int
		son       string
	)
	fmt.Println(king)
	fmt.Println(bin)
	fmt.Println(son)
}
