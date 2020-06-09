package main

func main() {
	input := []int{0, 1, 1, 1, 1}
	output := HasSingleCycle(input)
	println(input, output)
}

func HasSingleCycle(array []int) bool {

	point, elementsVisited := 0, 0

	for elementsVisited < len(array) {

		if elementsVisited > 0 && point == 0 {
			return false
		}

		elementsVisited++

		move := array[point]

		point = (point + move) % len(array)

		if point < 0 {
			point = len(array) + point
		}
	}

	return point == 0 && elementsVisited > 0
}
