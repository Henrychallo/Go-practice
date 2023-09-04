// The select statement lets a goroutine wait on multiple communication operations(send&receive)
// A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready.
package main

import "fmt"

func fib(c, quit chan int) { // fibonacci
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

// main function in channels.go file
