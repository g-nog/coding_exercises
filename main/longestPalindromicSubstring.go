package main

func main() {
	expected := "xyzzyx"
	output := LongestPalindromicSubstring("abaxyzzyxf")

	print(expected, output)
}

func LongestPalindromicSubstring(str string) string {

	longestLeft, longestRight := 0, 1

	for i, _ := range str {

		oddLeft, oddRight := getLongestPalindromic(str, i-1, i+1)
		evenLeft, evenRight := getLongestPalindromic(str, i-1, i)

		maxOdd := oddRight - oddLeft
		maxEven := evenRight - evenLeft
		maxCurrent := longestRight - longestLeft

		if maxOdd > maxCurrent {
			longestLeft = oddLeft
			longestRight = oddRight
		}

		if maxEven > maxCurrent {
			longestLeft = evenLeft
			longestRight = evenRight
		}
	}

	// Write your code here.
	return str[longestLeft:longestRight]
}

func getLongestPalindromic(s string, left int, right int) (int, int) {

	for left >= 0 && (right < len(s)) {
		println(string(s[left]), "->", string(s[right]))
		if s[left] != s[right] {
			break
		}
		left--
		right++
	}

	return left + 1, right
}
