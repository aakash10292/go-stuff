package main

import (
	"../AccessUnexported/accUnexp"
	"fmt"
)

func main() {
	counter := accUnexp.New(15)
	fmt.Printf("Counter: %d\n", counter)
	/*
	   Eventhough alertCounter was unexported, we got access to it
	   through New() which was exported
	*/
}
