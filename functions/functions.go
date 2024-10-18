package main

import "fmt"

func factorial(number int) int {
	if number < 0 {
		return 0
	}
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
}

func main() {
	var number int
	number = -1
	a := factorial(number)
	if a == 0 {
		fmt.Println("math: square root of negative number: ", number)
	} else {
		fmt.Println("Fatorial ", number, " = ", a)
	}
}
