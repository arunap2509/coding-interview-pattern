package slidingwindow

func FindAllAnagramOfPatternInString(input, pattern string) [] int{
	if len(input) < len(pattern) {
		return []int{}
	}

	fp := 0
	sp := 0
	patternMap := make(map[rune]int)
	matched := 0
	var index []int

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
			index = append(index, fp)
		}

		// Q - why are we doing this?
		// A - this will not trigger untill we process the string to the length of the pattern
		// so once we process the length then we go to the next window in the next loop
		// before that we need to make sure we remove the first element from the window and 
		// add that to the map, so that we can consider the index 2 -> new sp as the next window
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

	return index
}