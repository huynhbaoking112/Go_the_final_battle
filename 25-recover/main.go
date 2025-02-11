package main

import "fmt"

func testError() {
	defer func() {

		if i := recover(); i != nil {
			fmt.Println("error in test error ", i)

		}
	}()

	fmt.Println("King ne")

	panic("Some thing went wrong!")

}

func main() {
	testError()
	fmt.Println("King ne2")
}
