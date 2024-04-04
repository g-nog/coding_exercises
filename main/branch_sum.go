package main

type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func (tree *BinaryTree) sumx(sum int, array *[]int) {

	currSum := sum + tree.Value

	if tree.Left != nil {
		tree.Left.sumx(currSum, array)
	}

	if tree.Right != nil {
		tree.Right.sumx(currSum, array)
	}

	if tree.Right == nil && tree.Left == nil {
		*array = append(*array, currSum)
	}
}

func BranchSums(root *BinaryTree) []int {

	var array = make([]int, 0)

	root.sumx(0, &array)
	// Write your code here.
	return nil
}

func main() {
	root := &BinaryTree{
		Value: 1,
		Left: &BinaryTree{
			Value: 2,
			Left: &BinaryTree{
				Value: 4,
				Left: &BinaryTree{
					Value: 8,
				},
				Right: &BinaryTree{
					Value: 9,
				},
			},
			Right: &BinaryTree{
				Value: 5,
				Left: &BinaryTree{
					Value: 10,
				},
			},
		},
		Right: &BinaryTree{
			Value: 3,
			Left: &BinaryTree{
				Value: 6,
			},
			Right: &BinaryTree{
				Value: 7,
			},
		},
	}

	res := BranchSums(root)

	print(res)

}
