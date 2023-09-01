package main

import (
	"fmt"
)

func main() {
	fmt.Println("Start of main function")

	// Deferring a function call
	defer fmt.Println("Deferred function call")

	fmt.Println("End of main function")

	// defer statement is used to schedule a function call to be executed when the surrounding function returns.

	//stacking defers
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
