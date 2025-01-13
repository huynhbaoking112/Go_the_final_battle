package main

import (
	f "fmt"
	r "runtime"
)

func main() {
	f.Println(r.NumCPU() + 1)
}
