package main

import "fmt"

func main() {
	for a := 0; a < 10; a++ {
		fmt.Print(a, "\n")
	}

	fmt.Print("Complete test 1\n\n")
	for a := 0; a < 3; a++ {
		for b := 3; b > 0; b-- {
			fmt.Print(a, " ", b, "\n")
		}
	}
	fmt.Print("Complete test 1\n\n")

	sum := 1
	for sum < 100 {
		sum += sum
		fmt.Print("Sum -> ", sum, "\n")
	}
	fmt.Print("Complete test 1\n\n")
}
