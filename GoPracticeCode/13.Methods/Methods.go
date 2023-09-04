//Methods in Go allow you to define behaviors and operations that can be performed on instances of that type

package main

import (
	"fmt"
	"math"
)

// Define a struct type called "Circle" to represent a circle
type Circle struct {
	Radius float64
}

// This method calculates and returns the area of the circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// This method calculates and returns the circumference(perimeter) of the circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	// Create an instance of the Circle struct
	myCircle := Circle{Radius: 5.0}

	// Call the "Area" method on the Circle instance
	area := myCircle.Area()
	fmt.Println("Area of the circle:", area)

	// Call the "Perimeter" method on the Circle instance
	perimeter := myCircle.Perimeter()
	fmt.Println("Circumference of the circle:", perimeter)
}
