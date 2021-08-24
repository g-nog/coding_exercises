package main

import "fmt"

func main() {

	fmt.Printf("result: %v", BubbleSort([]int{8, 5, 2, 9, 5, 6, 3}))

}

func BubbleSort(array []int) []int {

	end := len(array) - 1

	for x := 0; x < len(array)-1; x++ {

		swaps := 0

		for i := 0; i < end; i++ {

			if array[i] > array[i+1] {
				array[i] = array[i] + array[i+1]
				array[i+1] = array[i] - array[i+1]
				array[i] = array[i] - array[i+1]
				swaps++
			} else if array[i] < array[i+1] {
				swaps++
			}

			if i == end-1 {
				end--
			}

		}

		if swaps == 0 {
			break
		}
	}

	// Write your code here.
	return array
}
