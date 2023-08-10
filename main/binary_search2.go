package main

import (
	"fmt"
)

func main() {
	input := []int{2, 4, 6, 8}

	lookup := int(8)

	result := BinarySearch(input, lookup)
	fmt.Printf("%v", result)

}

func BinarySearch(array []int, target int) int {
	left, right := 0, len(array)-1
	for {
		m := int((left + right) / 2)
		if target == array[m] {
			return m
		} else if left == right {
			return -1
		}
		if target > array[m] {
			left = m + 1
		} else {
			right = m - 1
		}
	}
}

func BinarySearchRecursive(array []int, target int) int {
	// Write your code here.
	return -1
}
