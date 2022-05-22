package main

import "fmt"

func main() {
	nums := []int{1, 3, 5, 6}
	target := 5
	fmt.Println(searchInsert(nums, target))

	nums = []int{1, 3, 5, 6}
	target = 2
	fmt.Println(searchInsert(nums, target))

	nums = []int{1, 3, 5, 6}
	target = 7
	fmt.Println(searchInsert(nums, target))
}

// binary search
func searchInsert(nums []int, target int) int {
	var left int = 0
	var right int = len(nums) - 1

	for left <= right {
		var mid int = (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
