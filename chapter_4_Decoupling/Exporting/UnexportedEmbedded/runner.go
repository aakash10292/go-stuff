// Sample program to show how to create values from exported types with
// embedded unexported types.
package main

import (
	"fmt"

	"../UnexportedEmbedded/users"
)

func main() {

	// Create a value of type Manager from the users package.
	u := users.Manager{
		Title: "Dev Manager",
	}

	// Set the exported fields from the unexported user inner type.
	u.Name = "Chole"
	u.ID = 10

	fmt.Printf("User: %#v\n", u)
}

