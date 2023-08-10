package main

import (
	"errors"
	"fmt"
	"sort"
)

/*
*
Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.

Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
*/
func main() {

	nums1 := []int{
		1, 9,
	}

	nums2 := []int{
		2, 4,
	}

	//1,2,3,4
	result, _ := avg(nums1, nums2)

	fmt.Printf("%f\n", result)
}

// θ(n log(n))
func avg(arr1 intArrayX, arr2 intArrayX) (float64, error) {

	if arr1.isEmptyOrNil() && arr2.isEmptyOrNil() {
		return 0, errors.New("both inputs are empty")
	}

	mergedArray := append(arr1, arr2...)

	sort.Ints(mergedArray)

	size := len(mergedArray)

	if size%2 == 0 {
		//pair
		middle := size / 2
		middle1 := (size / 2) - 1

		x := mergedArray[middle1 : middle+1]

		sum := 0.0

		for i := range x {
			sum += float64(x[i])
		}

		return sum / 2, nil

	} else {
		//odd
		middle := (size - 1) / 2
		return float64(mergedArray[middle]), nil
	}
}

type intArrayX []int

func (i intArrayX) isEmptyOrNil() bool {

	if i == nil {
		return true
	}

	if len(i) == 0 {
		return true
	}

	return false
}
