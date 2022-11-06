package slidingwindow

import "math"

func MinimumWindowSubsequence(str1, str2 string) string {
	indexS1 := 0
	indexS2 := 0
	minLength := math.MaxInt64
	minimumSubSequence := ""

	for indexS1 < len(str1) {
		if str1[indexS1] == str2[indexS2] {
			indexS2++

			if indexS2 == len(str2) {

				start, end := indexS1, indexS1 + 1

				indexS2--

				for indexS2 >= 0 {

					if str1[start] == str2[indexS2] {
						indexS2--
					}

					start--
				}

				// this will bring the start to the correct index where the current substring begins
				start++

				if end - start < minLength {
					minLength = end - start
					minimumSubSequence = str1[start:end]
				}

				// this will assign the indexs1 to the correct substring index as start
				indexS1 = start
				indexS2 = 0
			}
		}

		// we have to process from the next index as thats the new window 
		indexS1++
	}

	return minimumSubSequence
}