// In Go, an interface is a type that specifies a set of methods.
// A struct (or any custom type) can then implement these methods to satisfy the interface.
package main

import (
	"fmt"
	"math"
)

// A Shape interface with two method signatures: Area and Perimeter.
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Define a Circle type that implements the Shape interface.
type Circle struct {
	Radius float64
}

// Implement the Area method for Circle.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Implement the Perimeter method for Circle.
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Define a Rectangle type that implements the Shape interface.
type Rectangle struct {
	Width  float64
	Height float64
}

// Implement the Area method for Rectangle.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Implement the Perimeter method for Rectangle.
func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

func main() {
	// Create a Circle and a Rectangle.
	circle := Circle{Radius: 5}
	rectangle := Rectangle{Width: 4, Height: 6}

	// Calculate and print the areas and perimeters of these shapes.
	shapes := []Shape{circle, rectangle}

	for _, shape := range shapes {
		fmt.Println("Shape:", shape)
		fmt.Println("Area:", shape.Area())
		fmt.Println("Perimeter:", shape.Perimeter())
		fmt.Println()
	}
}
