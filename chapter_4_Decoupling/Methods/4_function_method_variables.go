package main

import "fmt"

// User defined type, we use this to
// bind methods to
type data struct {
	name string
	age  int
}

// uses VALUE semantic receiver
func (d data) displayName() {
	fmt.Println("My name is ", d.name)
}

// uses POINTER semantic receiver
func (d *data) setAge(age int) {
	d.age = age
	fmt.Println(d.name, " is ", d.age, " years old")
}

func main() {
	d := data{
		name: "Bill",
	}

	fmt.Println("Proper Calls to Methods:")
	// How we actually call methods in go
	d.displayName()
	d.setAge(45)

	fmt.Println("\nWhat the Compiler is doing")

	// What Go is doing underneath
	data.displayName(d)
	(*data).setAge(&d, 45)

	// =========================================================================

	fmt.Println("\nCall Value Receiver Methods with Variable:")

	// Declare a function variable for the method bound to the d variable.
	// The function variable will get its own copy of d because the method
	// is using a value receiver.
	f1 := d.displayName

	// Call the method via the variable.
	f1()

	// Change the value of d.
	d.name = "Joan"

	// Call the method via the variable. We don't see the change.
	f1()

	// =========================================================================

	fmt.Println("\nCall Pointer Receiver Method with Variable:")

	// Declare a function variable for the method bound to the d variable.
	// The function variable will get the address of d because the method
	// is using a pointer receiver.
	f2 := d.setAge

	// Call the method via the variable.
	f2(45)

	// Change the value of d.
	d.name = "Sammy"

	// Call the method via the variable. We see the change.
	f2(45)
}
