package main

func main() {
	expected := true
	output := BalancedBrackets(")[]}")

	print(expected, output)
}

func BalancedBrackets(s string) bool {

	stack := make([]string, 0)

	for _, i := range s {

		val := string(i)

		if isOpenBracket(val) {
			stack = append(stack, val)
		} else if len(stack) == 0 {
			return false
		} else {
			n := len(stack) - 1 // Top element
			if stack[n] == reverse(val) {
				stack = stack[:n] // Pop
			} else {
				return false
			}
		}

	}

	// Write your code here.
	return len(stack) == 0
}

func isOpenBracket(s string) bool {
	openBrackets := []string{"(", "{", "["}
	for _, bracket := range openBrackets {
		if s == bracket {
			return true
		}
	}
	return false
}

func reverse(s string) string {
	switch s {
	case "(":
		return ")"
	case ")":
		return "("
	case "[":
		return "]"
	case "]":
		return "["
	case "{":
		return "}"
	case "}":
		return "{"
	default:
		return ""
	}
}
