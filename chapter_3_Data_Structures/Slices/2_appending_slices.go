package main

import "fmt"

func main() {
	// Declare a nil slice of strings
	var data []string // We used var, hence zero-value init

	emptySlice := []string{} // Empty literal initialization
	fmt.Printf("Empty String: %+v\n", emptySlice)
	/*
	   Unlike the other types(built-in and user-defined),
	   Empty literal init does NOT do zero value init.
	   There is a special struct available at run-time - Empty Slice
	   The pointer of this slice is not null
	   Use case - imagine a db call, you did a select, and returned empty.
	   In this case we don't want to return a nil slice (we'd probably want
	   to return nil on an error). We should return an empty slice.
	   The pointer of an Empty slice points to an Empty Struct - zero-allocation type.
	   We can have a million Empty slices. All of them point to the same 8-byte fixed "Empty Struct"
	   available at runtime.
	*/

	lastCap := cap(data)
	// Append ~100k strings to the slice
	for record := 1; record <= 1e5; record++ {
		// Use the built-in function append to add to the slice
		value := fmt.Sprintf("Rec: %d", record)
		data = append(data, value)

		/*
		   Append call above is using VALUE SEMANTICS -> it gets it own copy
		   of the slice, and hence can perform mutations in isolations without
		   any sideffects
		*/

		// When the capacity of the slice changes,
		// display the changes
		if lastCap != cap(data) {
			// Calculate the percent of change
			capChg := float64(cap(data)-lastCap) / float64(lastCap) * 100

			// Save the new values for cacpacity
			lastCap = cap(data)

			// Display the results
			fmt.Printf("Addr[%p]\tIndex[%d]\t\tCap[%d - %2.f%%]\n",
				&data[0],
				record,
				cap(data),
				capChg)
		}
	}
}

/*
   In the code above, we know before hand that we're going to insert
   1e5 records. In such a case, it is better (performance wise),
   to use make and create a slice of a particular size
   i.e data := make([]string, 0, 1e5)

   If you create the data slice as above , and run the program again,
   you should not see any output. This is because we never ever have to make
   a copy of the backing array to re-initialize to a new larger array
*/
