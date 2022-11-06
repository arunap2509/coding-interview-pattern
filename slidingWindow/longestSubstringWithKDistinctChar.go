package slidingwindow

import "math"

func LongestSubstringWithKDistinctCharacters(str string, k int) (string, int) {
	var distinct = make(map[byte]int)
	var substring string
	var maxLength int = math.MinInt32

	fp := 0
	sp := 0

	for sp < len(str) {
		distinct[str[sp]] += 1
		sp++

		for len(distinct) > k {
			distinct[str[fp]] -= 1
			if distinct[str[fp]] == 0 {
				delete(distinct, str[fp])
			}
			fp++
		}

		if (sp-fp) > maxLength {
			substring = str[fp:sp]
			maxLength = sp - fp
		}
	}

	return substring, maxLength
}