package main

func main() {

	sockMerchant(9, []int32{9, 10, 20, 20, 10, 10, 30, 50, 10, 20})

}

// Complete the sockMerchant function below.
func sockMerchant(n int32, ar []int32) int32 {

	m := make(map[int32]int)

	for _, val := range ar {
		m[val]++
	}

	total := 0

	for _, i := range m {
		for i >= 2 {
			total++
			i = i - 2
		}
	}

	return int32(total)

}
