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

	fmt.Println("recursive :")
	printLeafNodesRecursive(&tree)

	fmt.Println()

	fmt.Println("iteravite :")
	printLeafNodesIterative(&tree)

}

func printLeafNodesRecursive(node *TreeNode) {
	if node == nil {
		return
	}

	fmt.Printf("%v ", node.Value)

	printLeafNodesRecursive(node.Left)
	printLeafNodesRecursive(node.Right)
}

// need stack algo (LIFO)
func printLeafNodesIterative(node *TreeNode) {
	if node == nil {
		return
	}

	var stack Stack
	stack.Push(*node)

	for !stack.IsEmpty() {
		*node, _ = stack.Pop()

		fmt.Printf("%v ", node.Value)

		if node.Left != nil {
			stack.Push(*node.Left)
		}

		if node.Right != nil {
			stack.Push(*node.Right)
		}

	}

}

// stack
type Stack []TreeNode

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(val TreeNode) {
	*s = append(*s, val)
}

func (s *Stack) Pop() (TreeNode, bool) {
	if s.IsEmpty() {
		return TreeNode{}, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}
