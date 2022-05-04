package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	res := mergedTwoLists(list1, list2)
	for res != nil {
		fmt.Println(res.Val)
		res = res.Next
	}
}

func mergedTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := new(ListNode)
	dummy.Val = -1

	ptr := dummy
	ptr1 := list1
	ptr2 := list2

	for ptr1 != nil && ptr2 != nil {
		if ptr1.Val < ptr2.Val {
			ptr.Next = ptr1
			ptr1 = ptr1.Next
		} else {
			ptr.Next = ptr2
			ptr2 = ptr2.Next
		}
		ptr = ptr.Next
	}

	if ptr1 != nil {
		ptr.Next = ptr1
	}

	if ptr2 != nil {
		ptr.Next = ptr2
	}

	return dummy.Next
}
