package main

// ---------------------------------------------------------
// EXERCISE: Declare with bits
//
//  1. Declare a few variables using the following types
//    int
//    int8
//    int16
//    int32
//    int64
//    float32
//    float64
//    complex64
//    complex128
//    bool
//    string
//    rune
//    byte
//
// 2. Observe their output
//
// 3. After you've done, check out the solution
//    and read the comments there
//
// EXPECTED OUTPUT
//  0 0 0 0 0 0 0 (0+0i) (0+0i) false 0 0
//  ""
// ---------------------------------------------------------

import "fmt"

func main() {
	var king1 int
	var king2 int8
	var king3 int16
	var king4 int32
	var king5 int64
	var king6 float32
	var king7 float64
	var king8 complex64
	var king9 complex128
	var bin bool
	var bin1 string
	var bin2 rune
	var bin3 byte

	fmt.Println(king1)
	fmt.Println(king2)
	fmt.Println(king3)
	fmt.Println(king4)
	fmt.Println(king5)
	fmt.Println("king6", king6)
	fmt.Println("king7", king7)
	fmt.Println("king8", king8)
	fmt.Println("king9", king9)
	fmt.Println("bin", bin)
	fmt.Println("bin1", bin1)
	fmt.Println("bin2", bin2)
	fmt.Println("bin3", bin3)
}
