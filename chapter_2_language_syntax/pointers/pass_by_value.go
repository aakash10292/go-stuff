/* Pass by value semantics:
    In golang, variables are passed by value just as they are in C-based languages.
    Every function executes in it's own stack frame and can access/modify data ONLY available within this stack frame
*/
package main

func main() {
    // count is a variable local to main
	count := 10
	println("count:\t Value of[", count, "]\t Addr of[", &count, "]")
    //count gets passed-by-value as a parameter to increment ( Just like in c/c++)
	increment(count)
    // Statement below should still print 10
	println("count:\t Value of[", count, "]\t Addr of[", &count, "]")
}

// Function that takes an int value , increments it once and displays the value
func increment(inc int) {
    //inc here is local to increment i.e modifications to this value are visible only within this method
	inc++
    // Should print 11 with an address different from that printed on lines 10 and 14
	println("count:\t Value of[", inc, "]\t Addr of[", &inc, "]")
}
