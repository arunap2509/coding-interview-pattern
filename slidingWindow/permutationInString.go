package slidingwindow

func PermutationInString(input, pattern string) bool {
	if len(input) < len(pattern) {
		return false
	}

	fp := 0
	sp := 0
	patternMap := make(map[rune]int)
	matched := 0

	for _, p := range pattern {
		patternMap[p] += 1
	}

	for sp < len(input) {
		if _, ok := patternMap[rune(input[sp])]; ok {
			patternMap[rune(input[sp])]--

			if patternMap[rune(input[sp])] == 0 {
				matched++
			}
		}

		if matched == len(patternMap) {
			return true
		}

		// the use for fp in here is to put the removed count back in the map
		// the condition will always keep the pattern length between sp and fp
		if sp >= len(pattern) - 1 {
			if _, ok := patternMap[rune(input[fp])]; ok {
				if patternMap[rune(input[fp])] == 0 {
					matched--
				}
				patternMap[rune(input[fp])]++
			}
			fp++
		}

		sp++
	}

	return false
}