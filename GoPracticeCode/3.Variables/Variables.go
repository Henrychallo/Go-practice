package main

import (
	"fmt"
)

// default variables
var c, python, java bool

// variables with initializers
var x, y int = 1, 2

// Type conversion
var z int = 19
var f float64 = float64(z)

// constant "const"
const Pi = 3.14

func main() {
	fmt.Println(c, python, java)

	fmt.Println(x, y, z, Pi)

	//short variable declaration without using "var" (only done inside a function)
	k := 3
	fmt.Println(k)
}
