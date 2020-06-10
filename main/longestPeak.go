package main

func main() {
	array := []int{1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3}

	actual := LongestPeak(array)

	print(actual)
}

func LongestPeak(array []int) int {
	// Write your code here.

	//peakStart, peakEnd, longestPeak := 0, 1, 0
	size := len(array)

	max := 0

	for i, val := range array {

		if i == 0 || i == size-1 {
			continue
		}

		if val > array[i-1] && val > array[i+1] {
			peakSize := getPeakFromTip(i, array)

			if peakSize > max {
				max = peakSize
			}
		}

	}

	return max
}

func getPeakFromTip(tip int, array []int) int {
	begin, end := tip-2, tip+2

	for begin >= 0 && array[begin] < array[begin+1] {
		begin--
	}
	for end < len(array) && array[end] < array[end-1] {
		end++
	}

	return (end - begin) - 1
}
