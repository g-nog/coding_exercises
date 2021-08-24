package main

import "fmt"

func main() {

	input := []int32{2, 4, 6, 8, 10, 11, 14}

	lookup := int32(2)

	if binarySearch(input, lookup) {
		fmt.Printf("%v is within %v", lookup, input)
	} else {
		fmt.Printf("%v not present in %v", lookup, input)
	}
}

func binarySearch(sortedArray []int32, lookup int32) bool {

	size := len(sortedArray)

	middle := size / 2

	itemMiddle := sortedArray[middle]

	slice := make([]int32, 0)

	if size == 1 && lookup != itemMiddle {
		return false
	}

	if lookup == itemMiddle {
		return true
	} else if lookup > itemMiddle {
		slice = append(slice, sortedArray[middle:size]...)
	} else {
		slice = append(slice, sortedArray[0:middle]...)
	}

	return binarySearch(slice, lookup)
}

//func assertSorted(array []int32) bool {
//
//
//
//	for i := range array {
//		if
//	}
//}
