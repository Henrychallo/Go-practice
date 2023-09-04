// Goroutines are lightweight concurrent threads of execution.
// You can create a goroutine by using the go keyword followed by a function or method call.

package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 1000; i++ {
		fmt.Printf("%d ", i)
		time.Sleep(1 * time.Millisecond)
	}
}

func printLetters() {
	for char := 'a'; char <= 'z'; char++ {
		fmt.Printf("%c ", char)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	// Start two goroutines concurrently
	go printNumbers()
	go printLetters()

	// Sleep/wait for a while to allow goroutines to finish
	time.Sleep(2 * time.Second)

	fmt.Println("\nMain function exiting...")
}
