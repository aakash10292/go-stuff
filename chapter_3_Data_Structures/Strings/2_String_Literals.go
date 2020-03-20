/*
    This program prints a string constant with a single character in three different ways:
    - plain text
    - ASCII-only quoted string (quoted verb %+q)
    - Individual bytes in hex

    To avoid any confusion, we work with a raw string enclosed by back quotes, so it can
    contain only literal text. Regular strings, enclosed by double quotes, can contain 
    escape sequences (refer 1_Intro.go) 
*/


package main 
import "fmt"

func main() {
    const placeOfInterest = `⌘``

    fmt.Printf("plain string: ")
    // below line should print ⌘`
    fmt.Printf("%s", placeOfInterest)
    fmt.Printf("\n")

    fmt.Printf("quoted string: ")
    // below line sohuld print \u2318
    // remember, %+q escapes out any non-printable and non-ASCII bytes ALL WHILE INTERPRETING UTF-8
    fmt.Printf("%+q", placeOfInterest)
    fmt.Printf("\n")

    // The loop below should print e2 8c 98.
    // what are these ? These bytes are the utf-8 encoding of the hex value 2318
    fmt.Printf("hex bytes: ")
    for i := 0; i < len(placeOfInterest); i++ {
        fmt.Printf("%x ", placeOfInterest[i])
    }
    fmt.Printf("\n")
}
