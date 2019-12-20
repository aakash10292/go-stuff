/*
   This program demonstrates how one can circumvent the undesrirable
   mutation side-effect while using slices of slice.
*/

package main

import "fmt"

func main() {
	slice1 := make([]string, 5, 8)
	slice1[0] = "Apple"
	slice1[1] = "Banana"
	slice1[2] = "Orange"
	slice1[3] = "Grape"
	slice1[4] = "Pear"
	inspectSlice(slice1)
	fmt.Println("*************************")

	fmt.Println("Slices with mutation")

	slice2 := slice1[2:4]
	inspectSlice(slice2)

	fmt.Println("*************************")
	slice2 = append(slice2, "Change from slice 2")

	inspectSlice(slice1)
	inspectSlice(slice2)
	fmt.Println("*************************")

	// Reset slice1 to it's original state
	slice1[4] = "Pear"
	fmt.Println("Slices without mutation")

	slice3 := slice1[2:4:4]
	inspectSlice(slice3)

	fmt.Println("*************************")
	slice3 = append(slice3, "Change from slice 3")

	inspectSlice(slice1)
	inspectSlice(slice3)
	fmt.Println("*************************")
}

// inspectSlice exposes the slice header for review.
func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i, s := range slice {
		fmt.Printf("[%d] %p %s\n",
			i,
			&slice[i],
			s)
	}
}
