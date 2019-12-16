/*
There are no stackoverflows in golang. Every go routine is allocated
2K bytes of stack space, which grows dynamically by 25% everytime
a limit is reached. This involves copying contents of existing stack,
mimicing pointers between stack frames (if any) etc, onto a new (larger)
stack.
*/

package main

// run with 10 and then run with 1024
const size = 10

func main() {
	s := "HELLO"
	stackCopy(&s, 0, [size]int{})
}

// Function StackCopy runs recursively increasing
// the size of the stack
func stackCopy(s *string, c int, a [size]int) {
	// when run with size 10, the below print statement
	// should print the same output.
	// However, when run with size 1024, we can see different
	// addresses for the string "HELLO" as the stack reaches it's
	// limit and gets reallocated
	println(c, s, *s)
	c++
	if c == 10 {
		return
	}
	stackCopy(s, c, a)
}

/* Food for thought :
Imagine we had a bunch of co-routines, each with it's own stack.
Imagine if somewhere in these stacks, we had pointers pointing to memory locations
of different stacks. Then for every relocation, we would have to keep track of inter-stack references,
make them point to the new location(after re-allocation). In terms of stop-the-world latency,
doing such a thing dynamically is as bad as it could get.
This is another case where escape analysis kicks in. The compiler would figure out that a piece of data
would be referred across go routine stacks, and thus constructs the value on the heap.
In summary, values are constructed on the heap when :
* when values are shared across go routine boundaries
* any value that can't stay on the frame because of integrity issues
* any value where we don't know the size at compile time
*/
