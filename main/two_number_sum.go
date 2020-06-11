package main

import "sort"

func main() {
	a := []int{
		3, 5, -4, 8, 11, 1, -1, 6,
	}
	print(TwoNumberSum(a, 10))
}

func TwoNumberSum(array []int, target int) []int {
	sort.Ints(array)
	result := make([]int, 0)

	left := 0
	right := len(array) - 1

	for left < right {
		currentSum := array[left] + array[right]

		if currentSum == target {
			result = append(result, array[left], array[right])
			return result
		} else if currentSum < target {
			left = left + 1
		} else if currentSum > target {
			right = right - 1
		}
	}

	return result
}
