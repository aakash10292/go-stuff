package main

import (
	"../Basic/counter"
	"fmt"
)

func main() {
	Counter := counter.AlertCounter(10)
	fmt.Printf("Counter: %d\n", Counter)
    /*
    open counter/counter.go and rename AlertCounter to alertCounter
    Compile, and run basic.go.
    You should see an error
    // ./basic.go:9: cannot refer to unexported name counters.alertCounter
	// ./basic.go:9: undefined: counters.alertCounter

    This is because identifiers that start with a LOWERCASE LETTER ARE UNEXPORTED.
    In this case, main cannot access as unexported identifier from counter package.
    Hence, the error. 
    */
}
