// A map maps keys to values
package main

import "fmt"

func main() {
	// Declare and initialize an empty map with string keys and int values.
	scores := make(map[string]int)

	// Add key-value pairs to the map.
	scores["Alice"] = 92
	scores["Bob"] = 87
	scores["Charlie"] = 78

	// Access values by key.
	fmt.Println("Alice's score:", scores["Alice"])
	fmt.Println("Bob's score:", scores["Bob"])
	fmt.Println("Charlie's score:", scores["Charlie"])

	// Check if a key exists in the map.
	_, aliceExists := scores["Alice"]
	_, davidExists := scores["David"]
	fmt.Println("Alice exists in the map:", aliceExists)
	fmt.Println("David exists in the map:", davidExists)

	// Delete a key-value pair from the map.
	delete(scores, "Charlie")

	// Iterating over a map using for loop.
	for name, score := range scores {
		fmt.Printf("%s's score: %d\n", name, score)
	}
}
