/*
So far, we've been using either Built-in types or user-defined types
golang has another class of types called reference types
These are all the slices, maps , channels, inerface values, functions.
They are DS that have a pointer. When any of these types are set to their zero
value they are considered to be nill
*/

package main

import "fmt"

func main() {
	// Create a slice with a length of 5 elements
	fruits := make([]string, 5) // we use make to create slice
	/*
	   Slice is a 3 words structure :
	   * Pointer to backing store
	   * Length of elements that can be accessed from Pointer position
	   * Capacity - number of elements that exist in backing store from that pointer position
	   Capacity can be greater than length
	*/
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	// You can't access an index of a slice beyond its length.
	// fruits[5] = "Runtime error"

	// Error: panic: runtime error: index out of range

	fmt.Println(fruits) // passed using VALUE SEMATICS
	// Println gets its own copy of the 3-word slice.
	// The backing store is shared
	inspectSlice(fruits)
}

func inspectSlice(slice []string) { // Notice how you don't pass a size between []
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i, s := range slice { // VALUE SEMANTICS for -> creates it's own 3 word slice
		fmt.Printf("[%d] %p %s\n", i, &slice[i], s)
	}
}
