package main

func main() {

	a := [][]int32{
		{-1, -1, 0, -9, -2, -2},
		{-2, -1, -6, -8, -2, -5},
		{-1, -1, -1, -2, -3, -4},
		{1, -9, -2, -4, -4, -5},
		{-7, -3, -3, -2, -9, -9},
		{-1, -3, -1, -2, -4, -5},
	}

	print(hourglassSum(a))

}

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) int32 {

	rows, columns := len(arr), len(arr[0])

	max := int32(-100)

	for i := 0; i < rows-2; i++ {
		for j := 0; j < columns-2; j++ {
			curr_max := arr[i][j] + arr[i][j+1] + arr[i][j+2] +
				arr[i+1][j+1] + arr[i+2][j] + arr[i+2][j+1] + arr[i+2][j+2]
			if curr_max > max {
				max = curr_max
			}
		}
	}

	return max
}
