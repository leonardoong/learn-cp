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

	fmt.Println(diameterOfBinaryTree(&tree))
}

var maxDiameter int

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter = 0
	dfs(root)

	return maxDiameter
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := dfs(root.Left)
	right := dfs(root.Right)

	maxDiameter = max(maxDiameter, left+right)

	return 1 + max(left, right)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
