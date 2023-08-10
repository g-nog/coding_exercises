package main

func main() {

	println(BalancedBrackets2("()[]{}{"))
}

var openCloseBrackets = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
}

var closeOpenBrackets = map[string]string{
	")": "(",
	"]": "[",
	"}": "{",
}

// BalancedBrackets2  ([]){}
func BalancedBrackets2(s string) bool {

	stack := make([]string, 0)

	for i := range s {
		v := string(s[i])
		if _, ok := openCloseBrackets[v]; ok {
			stack = append(stack, v) //push
		} else {
			stackSize := len(stack)
			if openBracket, ok := closeOpenBrackets[v]; ok {
				if stackSize == 0 {
					return false
				}
				stackSize -= 1
				if openBracket == stack[stackSize] {
					stack = stack[:stackSize]
				} else {
					return false
				}
			}
		}
	}
	// Write your code here.
	if len(stack) == 0 {
		return true
	}
	return false
}
