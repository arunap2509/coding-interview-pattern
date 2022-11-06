package slidingwindow

import "math"

func FindBestTimeToBuyStocks(data []int) int {

	if len(data) == 0 {
		return 0
	}

	profit, min := math.MinInt32, math.MaxInt32

	for i := 0; i < len(data); i++ {
		min = Min(min, data[i])
		profit = Max(profit, data[i]-min)
	}

	return profit
}