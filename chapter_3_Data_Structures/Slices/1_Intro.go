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

    fmt.Println(fruits) // slice passed as VALUE to print function
	// Println gets its own copy of the 3-word slice.
	// The backing store is shared
	inspectSliceVals(fruits)

    fmt.Printf("Original slice pointer: %p\n", &fruits)
    inspectSliceRefs(fruits)

    fmt.Println("Inspecting types")
    inspectSliceTypes(fruits)
}

func inspectSliceVals(slice []string) { // Notice how you don't pass a size between []
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i, s := range slice { // VALUE SEMANTICS for -> creates it's own 3 word slice
		fmt.Printf("[%d] %s %s\n", i, slice[i], s)
	}
}
func inspectSliceRefs(slice []string) { 
    // VALUE semantics being used, the function has it's own slice reference
    fmt.Printf("Slice used in inspectSliceRefs: %p\n",&slice)
	for i, s := range slice { // VALUE SEMANTICS for -> creates it's own copy. anonymous, no name (rvalue)
		fmt.Printf("[%d] %p %p\n", i, &slice, &s)  // slice here is the slice passed to the function, NOT the range for copy
	}
}

func inspectSliceTypes(slice []string){
    for i,s := range slice {
        fmt.Printf("iterator[%d]:%T slice:%T s:%T\n",i,i, slice, s) // you should see that slice is of type 'slice' and s is a 'string'
    }
}
