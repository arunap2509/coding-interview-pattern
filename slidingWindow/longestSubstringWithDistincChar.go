package slidingwindow

import (
	"math"
)

func LongestSubstringWithDistinctChar(str string) int {
	freq := make(map[byte]int)
	fp := 0
	sp := 0
	longest := math.MinInt32

	for sp < len(str) {
		charProcessed := str[sp]
		freq[charProcessed] += 1
		sp++

		for freq[charProcessed] > 1 {
			freq[str[fp]]--
			if freq[str[fp]] == 0 {
				delete(freq, str[fp])
			}
			fp++
		}

		longest = Max(longest, sp-fp)
	}

	return longest
}

func LongestSubstringWithDistinctChar2(str string) int {
	charIndex := make(map[byte]int)
	fp := 0
	sp := 0
	longest := math.MinInt32

	for sp < len(str) {
		if _, ok := charIndex[str[sp]]; ok {
			fp = Max(fp, charIndex[str[sp]] + 1)
		}
		charIndex[str[sp]] = sp
		sp++
		longest = Max(longest, sp - fp)
	}

	return longest
}