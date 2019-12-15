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

	// Declaring a variable of type example and init using a struct literal.
	// Remember we can use the := operator when we want to declare an init with a non-zero value.
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.1415,
	}

	//We can display individual field values using the . operator"
	fmt.Println("Flag ", e2.flag)
	fmt.Println("Counter ", e2.counter)
	fmt.Println("pi ", e2.pi)
}
