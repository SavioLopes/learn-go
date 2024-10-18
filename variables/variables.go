package main

import (
	"fmt"
	"reflect"
)

func main() {
  var text = "string"
  fmt.Println(reflect.TypeOf(text))

  var integerOne = 1
  fmt.Println(reflect.TypeOf(integerOne))
  
  var float = 1.1
  fmt.Println(reflect.TypeOf(float))
  
  var str string = "string"
  fmt.Println(reflect.TypeOf(str))

  var integerTwo int
  fmt.Println(integerTwo)
  fmt.Println(reflect.TypeOf(integerTwo))

  b:= "hello world"
  // b = 1
  println(b)
}