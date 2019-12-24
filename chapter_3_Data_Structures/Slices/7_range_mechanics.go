package main

import "fmt"

func main() {
	friends := []string{"Annie", "Betty", "Charlie", "Doug", "Edward"}
	// Value sematic form of the for range
	for _, v := range friends {
		friends = friends[:2]
		fmt.Printf("v[%s]\n", v)
	}
	fmt.Println("************************")
	// Using the pointer semantic form of the for range.
	friends = []string{"Annie", "Betty", "Charley", "Doug", "Edward"}
	for i := range friends {
		friends = friends[:2]
		fmt.Printf("v[%s]\n", friends[i])
	}
	/*
	   While using pointer semantics, the code panics.
	   The for range took the length of the slice before iterating,
	   but during the loop that length changed. Now on the third iteration,
	   the loop attempts to access an element that is no longer
	   associated with the slice's length
	*/
}
