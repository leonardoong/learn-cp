package main

import (
	"fmt"
)

func main() {
	fmt.Println(canJump([]int{2, 3, 1, 1, 4}))
	fmt.Println(canJump([]int{3, 2, 1, 0, 4}))
}

func canJump(nums []int) bool {
	goal := len(nums) - 1

	for i := goal - 1; i >= 0; i-- {
		if i+nums[i] >= goal {
			goal = i
		}
	}

	if goal == 0 {
		return true
	}

	return false
}
