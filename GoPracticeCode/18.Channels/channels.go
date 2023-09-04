// Channels provide a way for goroutines to send and receive data with the channel operator, <-.
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c) // 0 to length/2
	go sum(s[len(s)/2:], c) // length/2 to upper bound
	x, y := <-c, <-c        // receive from c

	fmt.Println(x, y, x+y)

	//
	//
	//
	//
	//

	// main function for range&close.go
	fmt.Println("Output for 'range&close.go'")
	c = make(chan int, 10) //buffered channel
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}

	// main function for select.go
	fmt.Println("Output for 'select.go'")
	c = make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fib(c, quit)
}
