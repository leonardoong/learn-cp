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
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}

	res := removeNthFromEnd(h1, 3)
	arrRes := []int{}
	for res != nil {
		arrRes = append(arrRes, res.Val)
		res = res.Next
	}
	fmt.Println(arrRes)

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}

	left := dummy
	right := head

	for n > 0 {
		right = right.Next
		n--
	}

	for right != nil {
		left = left.Next
		right = right.Next
	}

	left.Next = left.Next.Next

	return dummy.Next
}
