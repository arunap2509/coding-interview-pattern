package slidingwindow

func WordConcatenation(str string, words []string) []int{
	var (
		wordFreqMap = make(map[string]int)
		resultIndices []int
		wordsCount = len(words)
		wordLength = len(words[0])
	)

	for _, word := range words {
		wordFreqMap[word]++
	}

	for i := 0; i <= len(str) - wordsCount * wordLength; i++ {
		wordsSeen := make(map[string]int)
		for j := 0; j < wordsCount; j++ {
			nextWordIndex := i + j * wordLength
			word := str[nextWordIndex:nextWordIndex+wordLength]

			if _, ok := wordFreqMap[word]; !ok {
				break
			}

			wordsSeen[word]++

			if wordsSeen[word] > wordFreqMap[word] {
				break
			}

			if j + 1 == wordsCount {
				resultIndices = append(resultIndices, i)
			}
		}
	}

	return resultIndices
}