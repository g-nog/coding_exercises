package main

func main() {
	println(countingValleys(8, "UDDDUDUU"))
}

// Complete the countingValleys function below.
func countingValleys(n int32, s string) int32 {


	result, vales := 0, 0

	for _, val := range s {
		if string(val) == "D" {
			result--
		} else if string(val) == "U" {
			result++
			if result == 0 {
				vales++
			}
		}
	}

	return int32(vales)

}
