package main

import "fmt"

//example is a type with different fields
type example struct {
	flag    bool
	counter int16
	pi      float64
} // no semi-colon needed

func main() {
	var e1 example // we used var. So, obviously we are declaring a variable of type example and setting it to it's zero value
	//Displaying the value now
	fmt.Printf("%+v\n", e1)
}
