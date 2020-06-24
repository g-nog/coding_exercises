package main

func main() {

	print(goodSegment([]int32{7, 48, 14, 33, 11, 17, 50, 25}, 12, 38))
}

/*
 * Complete the 'goodSegment' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY badNumbers
 *  2. INTEGER lower
 *  3. INTEGER upper
 */

func goodSegment(badNumbers []int32, lower int32, upper int32) int32 {

	m := make(map[int32]bool)

	for _, number := range badNumbers {
		m[number] = true
	}

	start, longestSegment := lower, int32(0)

	for x := lower; x <= upper; x++ { //linear time

		segSize := int32(0)

		if _, in := m[x]; in {

			segSize = x - start

			start = x + 1
		} else if x == upper {

			segSize = x - start + 1
		}

		if segSize > longestSegment {
			longestSegment = segSize
		}
	}

	return longestSegment
}
