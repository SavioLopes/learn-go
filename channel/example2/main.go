package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan int)

	go func() {
		for i := range 10 {
			canal <- i + 10
			time.Sleep(time.Second)
		}
	}()

	for v := range canal {
		fmt.Println(v)
	}
}
