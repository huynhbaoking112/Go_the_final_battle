package main

import "fmt"

func main() {

	var b byte
	b = 0

	fmt.Printf("%08b=%d\n", b, b)

	b = 255
	fmt.Printf("%08b=%d\n", b, b)

}
