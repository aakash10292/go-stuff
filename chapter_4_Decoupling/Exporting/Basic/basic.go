package main

import (
	"../Basic/counter"
	"fmt"
)

func main() {
	Counter := counter.AlertCounter(10)
	fmt.Printf("Counter: %d\n", Counter)
}
