package main

import "fmt"

func main() {
	// Create a mpa to track scores for players in a game
	scores := make(map[string]int)

	// Read the element at the key "anna". It is absent.
	// so we get the zero-value for this map's value type
	// Does this make sense ?
	score := scores["anna"]

	fmt.Println("Score:", score)

	/*
	   If we need to check the presence of a key we use
	   a 2-variable assignment. The 2nd variable is a
	   bool
	*/
	score, ok := scores["anna"]

	fmt.Println("Score:", score, "Present: ", ok)
}
