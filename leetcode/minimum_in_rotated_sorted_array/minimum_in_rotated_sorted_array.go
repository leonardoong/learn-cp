package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMin([]int{4, 5, 6, 7, 0, 1, 2}))
	fmt.Println(findMin([]int{5, 1, 2, 3, 4}))
	fmt.Println(findMin([]int{11, 13, 15, 17}))
}

func findMin(nums []int) int {
	l, r := 0, len(nums)-1
	min := nums[0]
	for l <= r {
		m := (l + r) / 2

		if nums[m] < min {
			min = nums[m]
		}

		if nums[m] >= nums[l] {
			l = m + 1
		} else {
			r = m - 1
		}
	}

	return min
}
