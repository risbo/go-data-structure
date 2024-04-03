package main

import (
	"fmt"
)

// import "os"

func main_b() {
	fmt.Print("Hello, World\n")

	var index int
	var price float64
	var active bool
	var codeNumber string

	fmt.Print("%T \t %T \t %T \t %T \n", index, price, active, codeNumber)
	fmt.Print("\n")
	fmt.Print("\n")
	fmt.Print("%v \t %v \t %v \t %q \n", index, price, active, codeNumber)
}
