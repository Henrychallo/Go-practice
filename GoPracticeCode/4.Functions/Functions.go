package main

import "fmt"

func multiply(x, y int) int {
	return x * y
}

//Functions returning multiple results
func swap(x, y string) (string, string) {
	return y, x
}

// "return" statement without arguments returns the named return values
func find(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(multiply(10, 4))

	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
