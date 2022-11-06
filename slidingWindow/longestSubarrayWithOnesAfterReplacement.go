package slidingwindow

func LongestSubarrayWithOnesAfterReplacement(arr []int, k int) int {
	fp := 0
	sp := 0
	maxLength := 0
	maxOnesRepeating := 0

	for sp < len(arr) {
		
		if arr[sp] == 1 {
			maxOnesRepeating++
		}

		sp++

		if (sp - fp - maxOnesRepeating) > k {
			if arr[fp] == 1 {
				maxOnesRepeating--
			}
			fp++
		}

		maxLength = Max(maxLength, sp - fp)
	}

	return maxLength
}