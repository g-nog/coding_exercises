package main

func main() {

	a := []int{5, 1, 22, 25, 6, -1, 8, 10}
	seq := []int{1, 6, -1, 10}

	print(IsValidSubsequence(a, seq))
}

func IsValidSubsequence(array []int, sequence []int) bool {

	seqIndex := 0

	for _, val := range array {

		if val == sequence[seqIndex] {
			seqIndex++
		}

		if seqIndex >= len(sequence) {
			break
		}
	}

	// Write your code here.
	return seqIndex == len(sequence)
}
