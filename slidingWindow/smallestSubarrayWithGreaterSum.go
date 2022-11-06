package slidingwindow

import "math"

func SmallestSubarrayWithGreaterSumThanK(arr []int, k int) (int, []int) {
	var subarrayLength int = math.MaxInt32
	var elem []int
	var changed bool

	fp := 0
	sp := 0
	sum := 0

	for sp < len(arr) {
		sum += arr[sp]
		sp++

		for sum >= k {
			subarrayLength, changed = hasMinChanged(subarrayLength, sp - fp)

			if changed {
				elem = arr[fp:sp]
			}

			sum -= arr[fp]
			fp++
		}
	}

	if subarrayLength == math.MaxInt32 {
		subarrayLength = 0
	}

	return subarrayLength, elem
}

func hasMinChanged(old, new int) (int, bool) {
	if old > new {
		return new, true
	}

	return old, false
}