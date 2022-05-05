package main

import "fmt"

func main() {
	x := []int{1, 1, 2}
	fmt.Println(removeDuplicates(x))

	x2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(x2))
}

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	var ptr1, ptr2 int = 0, 1
	for ptr1 < len(nums) {
		if nums[ptr1] != nums[ptr2] {
			ptr1++
			nums[ptr1] = nums[ptr2]
		}

		if nums[ptr1] == nums[len(nums)-1] {
			return ptr1 + 1
		}
		ptr2++
	}
	return ptr1 + 1
}
