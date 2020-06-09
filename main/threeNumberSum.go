package main

import "sort"

func main() {
	expected := [][]int{{-8, 2, 6}, {-8, 3, 5}, {-6, 1, 5}}
	output := ThreeNumberSum([]int{12, 3, 1, 2, -6, 5, -8, 6}, 0)

	print(expected)
	print(output)
}

func ThreeNumberSum(array []int, target int) [][]int {
	r := make([][]int, 0)
	sort.Ints(array)

	for i, val := range array {

		left, right := i+1, len(array)-1

		for left < right {
			leftVal := array[left]
			rightVal := array[right]

			possibleSum := val + leftVal + rightVal

			if possibleSum == target {
				r = append(r, []int{val, leftVal, rightVal})
				left++
				right--
			} else if possibleSum < target {
				left++
			} else if possibleSum > target {
				right--
			}
		}

	}

	// Write your code here.
	return r
}
