package main

func main() {
	println(jumpingOnClouds([]int32{0, 0, 1, 0, 0, 1, 0}))
}

// Complete the jumpingOnClouds function below.
func jumpingOnClouds(c []int32) int32 {

	jumps, _, pos := 0, 0, 0

	for pos < len(c) {
		if pos+1 < len(c) {
			if c[pos+1] == 0 {
				jumps++
			}
		}
		if pos+2 < len(c) {
			if c[pos+2] == 0 {
				jumps++
				pos++
			}
		}

		pos++
	}

	return int32(jumps)
}
