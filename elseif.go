package main

import "fmt"

func mainTest() {
	//readFromTerminal()
	//checkingRangeNumber()
	//validNumber()
	switchScan()
}

func switchScan() {
	code := 100
	switch code {
	case 200:
		fmt.Print("was 200\n")
		fallthrough
	case 100:
		fmt.Print("was 100\n")
		fallthrough

	default:
		fmt.Print("Complted test\n")

	}
}
func readFromTerminal() {
	fmt.Print("Enter a number: ")
	var input int
	fmt.Scanln(&input)
	validateNumber(input)

}
func validateNumber(number int) {

	if number%2 == 0 {
		fmt.Print("Is a even number\n")
	} else {
		fmt.Print("Is a odd\n")
	}

	fmt.Print("Value validated is: ", number)
}

func checkingRangeNumber() {
	fmt.Print("Enter a number:")
	var number int
	fmt.Scanln(&number)

	if number < 0 || number > 100 {
		fmt.Print("Number out of range for this test, must be 1..100")
	} else if number >= 0 && number < 50 {
		fmt.Print("Fail")
	} else if number >= 50 && number < 60 {
		fmt.Print("D Grade")
	} else if number >= 60 && number < 70 {
		fmt.Print("C Grade")
	} else if number >= 70 && number < 80 {
		println("B Grade")
	} else if number >= 80 && number < 10 {
		fmt.Print("A Grade")
	} else if number >= 90 && number <= 100 {
		fmt.Print("A+ Grade")
	}
}

func validNumber() {
	var x int = 10
	var y int = 20
	if x >= 10 {
		if x >= 10 {
			fmt.Print("Correct Number\n")
		}
	}

	fmt.Printf("Value of x = %d\n", x)
	fmt.Printf("Value of y = %d\n", y)

}
