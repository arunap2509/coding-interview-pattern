package slidingwindow

func SmallestWindowContainingSubstring(input, pattern string) string {
	if len(input) < len(pattern) {
		return ""
	}

	fp := 0
	sp := 0
	patternMap := make(map[rune]int)
	matched := 0
	changed := false
	minLength := len(input) + 1
	response := ""

	for _, p := range pattern {
		patternMap[p] += 1
	}

	for sp < len(input) {
		if _, ok := patternMap[rune(input[sp])]; ok {
			patternMap[rune(input[sp])]--
			if patternMap[rune(input[sp])] >= 0 {
				matched++
			}
		}

		for matched == len(pattern) {
			minLength, changed = hasMinChanged(minLength, sp - fp + 1)
			if changed {
				response = input[fp:sp+1]
			}

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

	return response
}