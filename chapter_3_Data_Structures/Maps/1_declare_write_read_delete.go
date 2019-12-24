/*
Example Program to show to initialize a map,
write to it, then read and delete from it.
*/
package main

import "fmt"

type user struct {
	name    string
	surname string
}

func main() {
	// Declate and make a map that stores values
	// of type user with a key of type string.
	users := make(map[string]user) // map[key_type]value_type

	// Add key/value pairs to the map, using VALUE SEMANTICS
	users["Roy"] = user{"Rob", "Roy"}
	users["Ford"] = user{"Henry", "Ford"}
	users["Mouse"] = user{"Mickey", "Mouse"}
	users["Jackson"] = user{"Michael", "Jackson"}

	// Read the value at a specific key
	mouse := users["Mouse"]

	fmt.Printf("%+v\n", mouse)

	// Replace the value at the Mouse key
	users["Mouse"] = user{"Jerry", "Mouse"}

	// Read the Mouse key again.
	fmt.Printf("%+v\n", users["Mouse"])

	// Delete the value at a specific key.
	delete(users, "Roy")

	// Check the length of the map. There should be only 3 elements
	fmt.Println(len(users))

	// It is safe to delete an absent key.
	delete(users, "Roy")
}
