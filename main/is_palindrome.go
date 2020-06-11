package main

func main() {
	expected := true
	output := IsPalindrome("abcdcba")

	print(expected, output)
}

func IsPalindrome(str string) bool {

	left, right := 0, len(str)-1

	for left < right {
		leftVal := string(str[left])
		rightVal := string(str[right])
		if leftVal == rightVal {
			left++
			right--
		} else {
			return false
		}
	}

	// Write your code here.
	return true
}
