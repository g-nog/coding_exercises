package main

func main() {
	expected := true
	output := IsPalindrome("abcdcba")

	print("Expected -> ", expected, "; Got-> ", output, "; Result -> ", func() string {
		if expected == output {
			return "SUCCESS"
		}

		return "FAIL"
	}())
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
