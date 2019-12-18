package main

import "fmt"

func main() {
	/*
	   * array declaration : variable_name [size]type
	   * Notice we are using var. recollect what that means ?
	     it means that the array is going to be zero-value initialized,
	     which further means that the 5 strings in the array below are going to
	     be zero-value initialized i.e  <null | 0>, where null is a pointer to a back array,
	     and 0 is the length of that array.
	*/
	var fruits [5]string
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"

	// Iterate over the array of strings.
	for i, fruit := range fruits {
		/*
		   range creates it's own array of fruits (value sematics)
		   fruit is, again, a copy of each one of these fruits
		*/
		fmt.Println(i, fruit) // a COPY of fruit (value sematics) is passed to Println()

		/*
				   In total, we have 5 copies of a fruit
		            * The actual string (residing in the data segment in this case)
		            * The copy in the fruits array
		            * The copy of the fruits array for range
		            * The fruit local variable in the for loop
		            * The fruit being passed to Println
		           BUT..
				   Important to note that all these copies SHARE THE SAME BACKING ARRAY( In this case, residing on the data segment)
		           Why does this matter? Imagine, the backing array was on the heap.
		           Then the garbage collector doesn't have to deal with 6 copies!! Since the others are on the stack (they clean up themselves)
		*/
	}
}
