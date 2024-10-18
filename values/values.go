package main

import "fmt"

func main() {
	fmt.Println("test " + "123")

	fmt.Println("2 == 2 =>", 2 == 2)

	fmt.Println("2 * -2 =>", 2*-2)

	fmt.Println("!!true =>", !!true)

	fmt.Println("2 != -2*-1 =>", 2 != -2*-1)

	fmt.Println("2*-1 !!= -2 =>", 2*-1 != -2)

	var x int = -2
	x *= -2
	fmt.Println("x=>", x)

	var y int = 2
	var z int = 2
	fmt.Println("y*z*2.0 =>", y*z*2.0)
	fmt.Println("y/z*2.0 =>", y/z*2.0)
	fmt.Println("x*2.0 =>", x*2.0)

	fmt.Println("5/2 =>", 5/2)
	fmt.Println("5%2 =>", 5%2)
	fmt.Println("5.0/2 =>", 5.0/2)
	fmt.Println("5/2.0 =>", 5/2.0)

	fmt.Println("true&&false||true&&2==2 =>", true && false || true && 2 == 2)
}
