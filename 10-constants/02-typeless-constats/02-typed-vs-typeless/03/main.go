package main

import "fmt"

func main() {
	const min = 42

	var f float64
	var s int

	f = min // OK when min is typeless
	s = min
	fmt.Println(f, s)
}
