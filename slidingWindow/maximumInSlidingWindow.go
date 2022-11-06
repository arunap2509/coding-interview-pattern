package slidingwindow

import "coding-interview-patterns/dataStructure"

func MaximumInSlidingWindow(arr []int, windowSize int) []int {
	result := make([]int, 0)

	window := datastructure.NewDeque()

	if len(arr) == 0 {
		return result
	}

	if windowSize > len(arr) {
		windowSize = len(arr)
	}

	for i := 0; i < windowSize; i++ {
		for !window.Empty() && arr[i] >= arr[window.Back()] {
			window.PopBack()
		}

		window.PushBack(i)
	}

	result = append(result, arr[window.Front()])

	for i := windowSize; i < len(arr); i++ {
		for !window.Empty() && arr[i] >= arr[window.Back()] {
			window.PopBack()
		}

		if !window.Empty() && window.Front() <= (i - windowSize) {
			window.PopFront()
		}

		window.PushBack(i)

		result = append(result, arr[window.Front()])
	}

	return result
}