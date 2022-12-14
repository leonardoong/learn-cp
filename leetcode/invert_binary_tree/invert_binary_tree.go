package main

import (
	"fmt"
)

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

	printLeafNodesRecursive(&tree)
	fmt.Println()
	printLeafNodesRecursive(invertTree(&tree))

}

// recursive
// TC : O(n)
// SC : O(n)
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	right := invertTree(root.Right)
	left := invertTree(root.Left)
	root.Left = right
	root.Right = left

	return root
}

func printLeafNodesRecursive(node *TreeNode) {
	if node == nil {
		return
	}

	fmt.Printf("%v ", node.Value)

	printLeafNodesRecursive(node.Left)
	printLeafNodesRecursive(node.Right)
}
