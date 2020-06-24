package main

// Do not edit the class below except for
// the insert, contains, and remove methods.
// Feel free to add new properties and methods
// to the class.
type BST struct {
	Value int

	Left  *BST
	Right *BST
}

// avg: O(log(n)) time | O(1) space
// worst: O(n) time | O(1) space
func (tree *BST) Insert(value int) *BST {

	if tree == nil {
		return nil
	}

	next := tree

	for true {

		if value >= next.Value {
			if next.Right == nil {
				next.Right = &BST{Value: value}
				break
			} else {
				next = next.Right
			}
		} else {
			if next.Left == nil {
				next.Left = &BST{Value: value}
				break
			} else {
				next = next.Left
			}
		}

	}

	// Write your code here.
	// Do not edit the return statement of this method.
	return tree
}

// avg: O(log(n)) time | O(1) space
// worst: O(n) time | O(1) space
func (tree *BST) Contains(value int) bool {

	currentNode := tree

	for currentNode != nil {
		if value == currentNode.Value {
			return true
		} else if value > currentNode.Value {
			currentNode = currentNode.Right
		} else { // lower than
			currentNode = currentNode.Left
		}
	}

	// Write your code here.
	return false
}

//WIP, remove has some issues
func (tree *BST) Remove(value int) *BST {

	tree.removeNode(value, nil)

	// Write your code here.
	// Do not edit the return statement of this method.
	return tree
}

func (tree *BST) removeNode(value int, parentNode *BST) {
	node := tree

	for node != nil {
		if value > node.Value {
			parentNode = node
			node = node.Right
		} else if value < node.Value {
			parentNode = node
			node = node.Right
		} else {
			if node.Right != nil && node.Left != nil {
				node.Value = node.Right.getMinValue()
				node.Right.Remove(node.Value)
			} else if parentNode == nil {
				if node.Left != nil {
					node.Value = node.Left.Value
					node.Right = node.Left.Right
					node.Left = node.Left.Left
				} else if node.Right != nil {
					node.Value = node.Right.Value
					node.Left = node.Right.Left
					node.Right = node.Right.Right
				} else {
					tree.Value = 0
				}
			} else if parentNode.Left == node {
				if node.Left == nil {
					parentNode.Left = node.Left
				} else {
					parentNode.Left = node.Right
				}
			} else if parentNode.Right == node {
				if node.Left == nil {
					parentNode.Right = node.Left
				} else {
					parentNode.Right = node.Right
				}
			}
			break
		}
	}
}

func (tree *BST) getMinValue() int {

	currentNode := tree
	lastMinVal := tree.Value

	for currentNode != nil {
		lastMinVal = currentNode.Value
		currentNode = currentNode.Left
	}

	return lastMinVal
}
