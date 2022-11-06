package slidingwindow

func AverageOfEachSubarrayOfKElements(arr []int, k int) []float32 {

	var avg []float32

	if k > len(arr) {
		return avg
	}

	firstPointer := 0
	secondPointer := 0
	sum := 0

	for secondPointer < len(arr) {
		sum += arr[secondPointer]
		secondPointer++

		if secondPointer - firstPointer == k {
			avg = append(avg, float32(sum)/float32(k))
			sum -= arr[firstPointer]
			firstPointer++
		}
	}

	return avg
}