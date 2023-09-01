// An array has a fixed size. A slice []T, on the other hand, is a dynamically-sized.
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13} //an array of 6 numbers

	var s []int = primes[1:4] //slice is formed by specifying two indices, a low and high bound(from 1 to 3)
	fmt.Println(s)            // a slice describes a part of an underlying array

	s = primes[:2] //slice default low bound is 0
	fmt.Println(s)

	s = primes[3:] //slice default upper bound is its length
	fmt.Println(s)

	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s) //slice length and capacity, len(s) and cap(s)

	a := make([]int, 5) //"make" function allocates a zeroed array and returns a slice that refers to that array, this is how you create dynamically-sized arrays
	fmt.Println("a", a)

	s = append(s, 2, 3, 4) //appending a slice s, results to a slice with elements of the original slice plus added values.
	fmt.Println(s)

	for i, v := range primes { //The range form of the for loop iterates over a slice (primes)
		fmt.Println("element", i, "is", v)
	}

	fmt.Println("elements are")
	for _, v := range primes { //You can skip the index or value by assigning to _
		fmt.Println(v)
	}
}
