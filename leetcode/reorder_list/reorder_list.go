package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	h1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
				},
			},
		},
	}
	h2 := &ListNode{
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

	res := reorderList(h1)
	arrRes := []int{}
	for res != nil {
		arrRes = append(arrRes, res.Val)
		res = res.Next
	}
	fmt.Println(arrRes)

	res2 := reorderList(h2)
	arrRes2 := []int{}
	for res2 != nil {
		arrRes2 = append(arrRes2, res2.Val)
		res2 = res2.Next
	}
	fmt.Println(arrRes2)
}

func reorderList(head *ListNode) *ListNode {
	// find middle node
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// middle node
	second := slow.Next
	slow.Next = nil

	// reverse second half of the linked list
	var prev *ListNode
	for second != nil {
		tmp := second.Next
		second.Next = prev
		prev = second
		second = tmp
	}

	// merge two linked list
	first, second := head, prev
	for second != nil {
		tmp1, tmp2 := first.Next, second.Next
		first.Next = second
		second.Next = tmp1
		first, second = tmp1, tmp2
	}

	return head
}
