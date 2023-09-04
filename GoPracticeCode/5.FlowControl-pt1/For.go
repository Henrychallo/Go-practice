package main

import "fmt"

func main() {

	sum := 0

	//For loop
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//"For" is used as "while" in Go
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	//
	//
	//
	//
	//
	//

	// main function for if.go file
	fmt.Println("Output for if.go")
	fmt.Println(sqrt(2), sqrt(-4))

	// main function for if&else.go file
	fmt.Println("Output for if&else.go")
	fmt.Println(pow(3, 2, 10), pow(3, 2, 1))
}
