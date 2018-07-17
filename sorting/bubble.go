package sorting

func BubbleSort(input []int) []int {
	lastIndex := len(input) - 1

	for i := lastIndex; i > 0; i-- {
		for j := 0; j < i; j++ {
			left := input[j]
			right := input[j+1]

			// Swap if left value > right value
			if left > right {
				leftCopy := input[j]
				input[j] = right
				input[j+1] = leftCopy
			}
		}
	}

	return input
}
