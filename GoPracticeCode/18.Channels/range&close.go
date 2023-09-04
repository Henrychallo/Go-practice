// A sender can close a channel to indicate that no more values will be sent.
package main

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

// main function in channels.go file
