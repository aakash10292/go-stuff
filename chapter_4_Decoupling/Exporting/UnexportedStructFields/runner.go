// Sample program to show how unexported fields from an exported struct
// type can't be accessed directly.
package main

import (
	"fmt"

	"../UnexportedStructFields/users"
)

func main() {

	// Create a value of type User from the users package.
	u := users.User{
		Name: "Chole",
		ID:   10,
		password: "xxxx",

	}
	// ./example4.go:21: unknown users.User field 'password' in struct literal

	fmt.Printf("User: %#v\n", u)
}

