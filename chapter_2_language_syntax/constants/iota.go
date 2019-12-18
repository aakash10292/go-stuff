package main

import "fmt"

func main() {
	const (
		A1 = iota // 0: Start at 0
		B1 = iota // 1: Increment by 1
		C1 = iota // 2: Increment by 1
		// We get this incrementing feature for free inside the constant block
	)

	fmt.Println("1:", A1, B1, C1)
	const (
		A2 = iota // 0 : Start at 0
		B2        // 1 : Increment by 1
		C2        // 2 : Increment by 1
		// We can just mention iota to initialize.
		// Don't have to repeat
	)

	fmt.Println("2:", A2, B2, C2)

	const (
		A3 = iota + 1 // 1 : Start at 0 + 1
		B3            // 2 : Increment by 1
		C3            // 3 : Increment by 1
		// We can have custom initialization with iota
		// Don't have to repeat
	)
	fmt.Println("3:", A3, B3, C3)

	const (
		A4 = 1 << iota // 1 : Shift 1 to left 0 times
		B4             // 2 : Shift 1 to left 1 time
		C4             // 4 : Shift 1 to left 2 times
		// We can have custom initialization with iota
		// Don't have to repeat
	)
	fmt.Println("4:", A4, B4, C4)
}
