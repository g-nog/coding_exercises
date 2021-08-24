package main

import (
	"fmt"
	"math/rand"
)

func main() {

	array := []int32{55, 90, 74, 20, 16, 46, 43, 59, 2, 99, 79, 10, 73, 1, 68, 56, 3, 87, 40, 78, 14, 18, 51, 24, 57, 89, 4, 62, 53, 23, 93, 41, 95, 84, 88}

	x := quickSort(array)

	fmt.Printf("%v", quickSort(x))
}

func quickSort(arr []int32) []int32 {

	size := len(arr)

	if size < 2 {
		return arr
	}

	if size == 2 {
		result := make([]int32, size)
		if arr[0] <= arr[1] {
			return arr
		}
		result[0] = arr[1]
		result[1] = arr[0]
		return result
	}

	pivot := arr[rand.Int()%size]

	left := make([]int32, 0)
	right := make([]int32, 0)

	for _, v := range arr {

		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		}
	}

	result := make([]int32, 0)

	result = append(quickSort(left), pivot)

	return append(result, quickSort(right)...)

}
