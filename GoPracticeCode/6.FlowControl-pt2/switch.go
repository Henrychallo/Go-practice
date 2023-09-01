//A switch statement is a shorter way to write a sequence of if - else statements.
//It runs the first case whose value is equal to the condition expression.

package main

import (
	"fmt"
	"time"
)

func main() {
	day := "Wednesday"

	switch day {
	case "Monday":
		fmt.Println("It's the beginning of the week.")
	case "Tuesday":
		fmt.Println("You're making progress.")
	case "Wednesday":
		fmt.Println("It's the middle of the week.")
	case "Thursday":
		fmt.Println("Almost there!")
	case "Friday":
		fmt.Println("Weekend is coming soon.")
	default:
		fmt.Println("Enjoy your weekend!")
	}

	//switch with no condition
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
