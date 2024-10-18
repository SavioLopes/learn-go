package main

import "fmt"

func main() {
	canal := make(chan int)

	go func() {
		canal <- 1 + 1
	}()

	fmt.Println(<-canal)
}
