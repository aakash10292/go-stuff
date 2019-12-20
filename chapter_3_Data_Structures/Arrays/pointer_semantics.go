package main

import "fmt"

func main() {
	friends := [5]string{"Annie", "Betty", "Charlie", "Doug", "Edward"}
	fmt.Printf("Bfr[%s] : ", friends[1])

	// POINTER SEMANTIC FORM OF for loop - we are only asking for the index
	// Here, UNLIKE VALUE SEMANTICS, the for range DOESN'T MAKE IT'S OWN COPY OF THE ORIGINAL ARRAY
	for i := range friends {
		friends[1] = "Jack" // Remember, this is POINTER SEMANTICS -> for-range DOESN'T work on it's own copy
		// -> we're modifying the original array
		if i == 1 {
			fmt.Printf("Aft[%s]\n", friends[1]) // Should print Jack
		}
	}

	fmt.Println("friends[1]: ", friends[1]) // Should still print Jack
}
