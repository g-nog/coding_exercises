package main

func main() {
	denoms := []int{1, 5, 10}
	n := 7
	println(MinNumberOfCoinsForChange(n, denoms))
}

// MinNumberOfCoinsForChange
// Given an array of positive integers representing coin denominations and a
// single non-negative integer representing a target amount of
// money, write a function that returns the smallest number of coins needed to
// make change for (to sum up to) that target amount using the given coin
// denominations.
// Note that you have access to an unlimited amount of coins. In other words, if
// the denominations are , you have access to an unlimited
// amount of
func MinNumberOfCoinsForChange(n int, denoms []int) int {
	dps := make([]int, n+1)
	dps[0] = 0
	const maxValue = 999999
	// setting up
	for i := range dps {
		if i == 0 {
			dps[i] = 0
		} else {
			dps[i] = maxValue
		}
	}
	i := 1
	for i <= n {
		for x := range denoms {
			c := denoms[x]
			if i-c >= 0 {
				dps[i] = minInt(dps[i], 1+dps[i-c])
			}
		}
		i++
	}
	// Write your code here.
	if dps[n] == maxValue {
		return -1
	} else {
		return dps[n]
	}
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
