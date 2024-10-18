package main

import (
	"fmt"
	"time"
)

func counter (count int) {
	for i := range count {
		fmt.Println(i)
		time.Sleep((time.Second))
	}
}

func main () {
	go counter(5)
	go counter(5)
	counter(5)
}