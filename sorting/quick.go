package sorting

func QuickSort(input []int, pivot int, start int, end int) []int {
	if end-start <= 0 {
		return input
	}

	// save the pivot value
	pivotValue := input[pivot]

	// initialize i and j
	i := start
	j := end

	// when i >= j, i is the partition index
	for i < j {
		// find the next right value that is < pivot
		// and move it to the left
		if input[j] < pivotValue {
			input[i] = input[j]
			// find the next left value that is >= pivot
			// and move it to the right
			for i < j && input[i] < pivotValue {
				i++
			}
			input[j] = input[i]
		}
		j--
	}

	// set partition point
	input[i] = pivotValue

	// sort left recursively
	input = QuickSort(input, 0, 0, i)

	// sort right recursively
	input = QuickSort(input, i+1, i+1, end)

	return input
}
