package slidingwindow

import "math"

func LongestSubstringHavingSameLetterAfterReplacement(str string, k int) int {
	fp := 0
	sp := 0
	freq := make(map[byte]int)
	maxLength := math.MinInt32
	maxRepeating := math.MinInt32

	for sp < len(str) {
		freq[str[sp]]++
		maxRepeating = Max(maxRepeating, freq[str[sp]])
		sp++

		if (sp - fp - maxRepeating) > k {
			freq[str[fp]]--
			fp++
		}

		maxLength = Max(maxLength, sp-fp)
	}
	return maxLength
}