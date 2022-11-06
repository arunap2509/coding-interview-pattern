package slidingwindow

import "math"

func MaximumSumSubarrayOfSizeK(arr []int, k int) int {
	if k > len(arr) {
		return 0
	}

	max := math.MinInt32
	fp := 0
	sp := 0
	sum := 0

	for sp < len(arr) {
		sum += arr[sp]
		sp++

		if sp - fp == k {
			max = Max(max, sum)
			sum -= arr[fp]
			fp++
		}
	}

	return max
}

func Max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}

	return a
}