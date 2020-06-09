package main

func main() {
	array := []int{5, 1, 22, 25, 6, -1, 8, 10}
	seq := []int{1, 6, -1, -1}

	print(IsValidSubsequence(array, seq))

}

func IsValidSubsequence(array []int, sequence []int) bool {

	seq_pos, seq_size := 0, len(sequence)

	for _, i := range array {
		if sequence[seq_pos] == i {
			seq_pos++
			if seq_pos >= seq_size {
				return true
			}
		}

	}
	// Write your code here.
	return false
}
