/*
    golang's standard library provides strong support for interpreting UTF-8 text.
    If the for range loops doesn't suit your purpose, it is very likely that you will
    find something useful in a package in the library. 

    Most important such package is the unicode/utf8. It containes helper routines to validate, 
    disassemble and reassemble UTF-8 strings. 
*/

package main

import (
    "fmt"
    "unicode/utf8"
)

func main() {
    const nihongo = "日本語"
    for i, w := 0, 0; i < len(nihongo); i += w {
        runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
        fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
        w = width
    }
    // The output should be the same as that obtain using a for range loop
}

