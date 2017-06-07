package bubbleSort

// Sort takes a slice of intgers and returns the slice sorted at O(n^2) time complexity
func Sort(values []int) []int {
	// temp makes it so we can swap values around
	var temp int

	// sorted gets set to true every main loop, but reverts to false if there is a sorting action.
	sorted := false

	// Main loop
	for !sorted {
		// sorted will be true if there are no swaps
		sorted = true

		// Loop through the array
		for i := 0; i < len(values)-1; i++ {

			// If the current value in the list is larger than the one after it, it should "bubble up" past the next element
			if values[i] > values[i+1] {
				// Because we had to make a change to the slice, we're not actually sorted yet
				sorted = false

				// Temporary variable to store our current value
				temp = values[i]

				// Move the next value to the current index, overwriting what was there
				values[i] = values[i+1]

				// Write the current value into the next index, overwriting the old, smaller value
				values[i+1] = temp
			}
		}
	}

	return values
}
