package main

import (
	"fmt"
)

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))

}

func search(nums []int, target int) int {
	idx := -1

	// binary search
	i := 0
	j := len(nums) - 1
	for i <= j {
		mid := (i + j) / 2
		if nums[mid] == target {
			return mid
		}

		if target < nums[mid] {
			j = mid - 1
		}

		if target > nums[mid] {
			i = mid + 1
		}
	}

	return idx
}
