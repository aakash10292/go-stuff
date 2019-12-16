/*
Construction of a value tells us nothing. In go, sharing is everything.
The go compiler performs static escape analysis to determine where
a value has to be created - stack or heap,
BASED ON HOW THE VALUE IS INTEDED TO BE SHARED

General Guideline :
Never use pointer semantics during construction IF
you're going to assign that value to a variable.
ONLY use pointer semantics during construction IF
* You're constructing on a return
* You're constructing inside a function call

Again, Construction doesn't tell you where something is in memory.
IT depends entirely on how it's shared
*/

package main

type user struct {
	name  string
	email string
}

func main() {
	// There are no constructors in go.
	// Instead we have something called "Factory Functions"
	u1 := createUserV1()
	// u1 is main() function's OWN COPY of a user value
	u2 := createUserV2()
	// u2 is

	println("u1", &u1, "u2", u2)
}

// Factory function
// Uses value semantics to return a User value
func createUserV1() user {
	// Literal construction of a variable `u` of value type `user`
	u := user{
		name:  "Aakash",
		email: "aakash@mail.edu",
	}
	// Sharing the user value with the print function
	println("V1", &u)

	// returns a copy of user i.e value sematics
	return u
}
// Factory function 
// Uses pointer semantics to return a pointer to a User value
func createUserV2() *user {
	u := user{
		name:  "Aakash",
		email: "aakash@mail.edu",
	}

	println("V2", &u)

	// If you're coming from a c/c++ background, this should raise all sorts of alarms.
	// * we're returning a pointer to a variable constructed on the stack
	// * Should result in undefined behavior/exception/seg fault
	return &u
    // If you run the program, you see that there's no error
    // Reason: 
    // Compiler performs escape analysis, realises that u2 is going to be shared up the stack (stack in system grows top to down)
    // Therefore, construction happens immediately on the heap, NOT ON THE STACK
    // Another interesting point - u represents a value of type user, (on heap, not on stack)
    // We know that if we want to access anything that is not in current stack frame, we can only do that using a pointer. HMMMMMMM
    // From a sytax perspective, u is a value on heap and go keeps it simple by allowing you to manipulate u value through VALUE SEMANTICS
    // But underneath, it is converted to pointer to allow access to the heap value. 
    // Syntax is abstracting away details of what's happening 
}



// See escape analysis and inling decisions by using the -gcflags options.
// build this file as : go build -gcflags "-m -m" escape_analysis.go
