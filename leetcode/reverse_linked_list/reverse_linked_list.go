package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// [1 2 3 4]
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
				},
			},
		},
	}

	// iterative solution
	res := reverseList(h1)
	arrRes := []int{}
	for res != nil {
		arrRes = append(arrRes, res.Val)
		res = res.Next
	}
	fmt.Println(arrRes)

	// recursive solution
	recRes := recReverseList(h2)
	recArrRes := []int{}
	for recRes != nil {
		recArrRes = append(recArrRes, recRes.Val)
		recRes = recRes.Next
	}
	fmt.Println(recArrRes)
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil {
		nxt := curr.Next
		curr.Next = prev
		prev = curr
		curr = nxt
	}

	return prev
}

func recReverseList(head *ListNode) *ListNode {

	if head == nil {
		return nil
	}

	newHead := head
	if head.Next != nil {
		newHead = recReverseList(head.Next)
		head.Next.Next = head
	}
	head.Next = nil

	return newHead
}
