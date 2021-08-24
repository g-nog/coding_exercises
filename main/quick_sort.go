package main

import "fmt"

func main() {

	x := quickSort([]int32{9, 10, 3, 0, 1, 8, 12})

	fmt.Printf("%v", quickSort(x))
}

func quickSort(arr []int32) []int32 {

	len := len(arr)

	if len < 2 {
		return arr
	}

	if len == 2 {
		result := make([]int32, len)
		if arr[0] <= arr[1] {
			return arr
		}
		result[0] = arr[1]
		result[1] = arr[0]
		return result
	}

	pivot := arr[len-1]
	left := make([]int32, 0)
	right := make([]int32, 0)

	for _, v := range arr {

		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		}
	}

	resultLeft := append(quickSort(left), pivot)
	resultRight := append(quickSort(right))

	return append(resultLeft, resultRight...)

}
