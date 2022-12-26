package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	noCycle := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	fmt.Println(hasCycle(noCycle))

	cycle := &ListNode{
		Val: 1,
	}

	list := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val:  5,
					Next: cycle,
				},
			},
		},
	}

	cycle.Next = list

	fmt.Println(hasCycle(cycle))
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false
}
