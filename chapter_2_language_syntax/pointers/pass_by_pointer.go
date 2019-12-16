/* Pass by value semantics:
   In golang, variables are passed by pointers just as they are in C-based languages.
   We pass the address of the data as the parameters.
   Essentially, this is pass by value - only difference is that the data being copied and passed is an address
   Pointer variables are used to store addresses.
*/
package main

func main() {
	// count is a variable local to main
	count := 10
	println("count:\t Value of[", count, "]\t Addr of[", &count, "]")
	// ADDRESS of count gets passed-by-value as a parameter to increment ( Just like in c/c++)
	increment(&count)
	// Statement below should still reflect the modification to count by above function call
	println("count:\t Value of[", count, "]\t Addr of[", &count, "]")
}

// Function that takes a POINTER to an int value and increments THE VALUE once
func increment(inc *int) {
	// inc here is local to increment BUT is a POINTER i.e modifications to the value POINTED to by this pointer persist
	// even after this method returns
	*inc++
	// Should print the address of count and another address (which is the address of the pointer itself)
	// Remember, pointers are data too - they have addresses just like other variables
	println("count:\t Value of[", inc, "]\t Addr of[", &inc, "]")
}
