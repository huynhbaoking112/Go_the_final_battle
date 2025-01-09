package main

//file scope
import "fmt"

//package scope
const ok = true

//package scope
func main() {
	var hello = "Hello"

	fmt.Println(hello, ok)
}
