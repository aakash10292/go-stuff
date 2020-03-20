/*
    A Rune is exactly the same as a codepoint. golang defines rune as an alias for INT32.
    This way programs can be clear when an integer value represents a codepoint.
    What one might think of as a character constant is called a rune constant in golang. 
    The type and value of the following expression:
    '⌘`'
    is rune with integer value 0x2318
*/


package main

import "fmt"

func main(){
    const nihongo = "日本語"
    /*
    A for range loop, unlike a regular for loop, decode one UTF-8 encoded rune on each iteration.
    Each time around the loop, the index of the loop is the starting position of the current rune,
    measured in bytes, and the code point is its value.
    */
    for index, runeValue := range nihongo {
        fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
        // %#U shows a codepoint's unicode value and its printed representation
    }
}
