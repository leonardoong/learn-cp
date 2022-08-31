package main

import (
	"fmt"
)

func main() {
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 0))
	fmt.Println(search([]int{4, 5, 6, 7, 0, 1, 2}, 3))
	fmt.Println(search([]int{1}, 1))
}

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	// use binary search
	for l <= r {
		mid := (l + r) / 2
		if target == nums[mid] {
			return mid
		}

		// left sorted portion
		if nums[l] <= nums[mid] {
			if target > nums[mid] || target < nums[l] {
				l = mid + 1
			} else {
				r = mid - 1
			}
			// right sorted portion
		} else {
			if target < nums[mid] || target > nums[r] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		}

	}

	return -1

}
