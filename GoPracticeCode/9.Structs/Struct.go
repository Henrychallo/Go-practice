// A struct is a collection of fields.
package main

import "fmt"

type Vertex struct {
	X, Y int //fields X and Y

}

func main() {
	v := Vertex{1, 2}

	fmt.Println(v.X) //struct fields can be accessed by using a dot
}
