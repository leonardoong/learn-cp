package main

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	tree := TreeNode{
		Value: 100,
		Left: &TreeNode{
			Value: 50,
			Left: &TreeNode{
				Value: 10,
			},
			Right: &TreeNode{
				Value: 25,
			},
		},
		Right: &TreeNode{
			Value: 90,
			Left: &TreeNode{
				Value: 60,
				Left: &TreeNode{
					Value: 55,
				},
			},
			Right: &TreeNode{
				Value: 80,
			},
		},
	}

	fmt.Println(maxDepth(&tree))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
