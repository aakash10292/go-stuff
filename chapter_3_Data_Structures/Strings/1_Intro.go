/*
    A string is effectively a READ-ONLY SLICE OF BYTES. A string holds arbitrary
    bytes. It is not required to hold unicode, UTF-8 or any other predefined format.
    IT IS SIMPLY A SEQUENCE OF BYTES. 
    
*/

package main

import "fmt"

func main() {
    // Sample is a string literal, \xNN is used to define a string constant
    const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

    fmt.Println("Println:")
    // This should produce ugly output as some of the bytes are not valid ASCII (nor valid UTF8)
    fmt.Println(sample)

    fmt.Println("Byte loop:")
    // To find out what the string really holds, we take the string
    // apart byte by byte and examine it. 
    // NOTE: indexing a string accesses individual bytes not characters
    for i := 0; i < len(sample); i++ {
        fmt.Printf("%x ", sample[i])
    }
    fmt.Printf("\n")

    // shorter way to do it - avoids the loop
    fmt.Println("Printf with %x:")
    fmt.Printf("%x\n", sample)
    // using a space between % and x produces a neater format
    fmt.Println("Printf with % x:")
    fmt.Printf("% x\n", sample)
    // %q(quoted) verb will escape any non-printable byte sequences.
    // This prevents output from being unambiguos. This technique is handy
    // when much of the string is intelligible but there are peculiarities to root out
    fmt.Println("Printf with %q:")
    fmt.Printf("%q\n", sample)
    // If strange symbols cause confusion(which they normally would), use + along with %q.
    // This escapes not only non printable sequences, but also any non-ASCII bytes, ALL
    // WHILE INTERPRETING UTF-8
    fmt.Println("Printf with %+q:")
    fmt.Printf("%+q\n", sample)
}


/*
These printing techniques are good to know when debugging the contents of strings.
All these methods work exactly the same with byte slices as they do for strings
*/
