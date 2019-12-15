package main

import "fmt"

func main() {
	// Declaring a variable e1 based on an anonymous literal struct type
	// We're using var so obviously initialisation it with zero-value
	var e1 struct {
		flag    bool
		counter int16
		pi      float32
	}

	fmt.Printf("%+v\n", e1)

	// Declaring a variable of an anonymous type and init with a struct literal
	e2 := struct {
		flag    bool //Notice , no commas during type declaration
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10, //Notice, commas used during value construction
		pi:      3.141592,
	}
	// A good use case for anonymous literal types : json unmarshalling, when we need type information only for this one place.
	// We wouldn't want to name something we'd be using in only one place (that would be pollution). Best to use above in such a case

	fmt.Printf("%+v\n", e2)

}
