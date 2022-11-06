package slidingwindow

import "math"

func FetchFuritsInBasket(fruits []string) (int, []string) {
	freq := make(map[string]int)

	fp := 0
	sp := 0
	var fruitsInBasket []string
	maxFruits := math.MinInt32

	for sp < len(fruits) {
		freq[fruits[sp]] += 1
		sp++

		for len(freq) > 2 {
			freq[fruits[fp]]--
			
			if freq[fruits[fp]] == 0 {
				delete(freq, fruits[fp])
			}
			fp++
		}
		fruitsInBasket = fruits[fp:sp]
		maxFruits = Max(maxFruits, sp-fp)
	}

	return maxFruits, fruitsInBasket
}