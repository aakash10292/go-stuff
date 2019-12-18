/*
Constants only exist at compile time.
golang specification says constants(OF A KIND) should be of atleast 256 bit precision
*/
package main

func main(){
    // Compiler can perform IMPLICIT CONVERSION of untyped constants
    // Untyped Constants.
    const ui = 12345    // kind: integer
    const uf = 3.141592 // kind: floating-point

    // Typed Constants still use the constant type system but
    // implicit conversion is restricted.
    const ti int = 12345        // type: int
    const tf float64 = 3.141592 // type: float64

    // Uncomment below line of code. It should throw a compile time error
    // as the number 1000 overflows uint8
    // const myUint8 uint8 = 1000
    
    // KIND PROMOTION : 
    /*
        Constant arithmetic supports different kinds.
        Kind Promotion is used to determine Kind in these scenarios
    */
    // variable answer will be of type float64
    var answer = 3 * 0.333 // KindFloat(3) * KindFloat(0.333)

    // Constant third will of KIND (not type) floating point
    const third = 1 / 3.0 // KindFloat(1) / KindFloat(3.0)

    // Constant zero will be of kind integer. 
    const zero = 1 / 3 // KindInt(1) / KindInt(3)

    // promotion happens to make sure we have like types to perform math
    const one int8 = 1
    const two = 2 * one // int8(2) * int8(1)
}
