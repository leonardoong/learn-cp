package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(maxSub([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
	fmt.Println(maxSub([]int{1}))
	fmt.Println(maxSub([]int{5, 4, -1, 7, 8}))
}

func maxSub(nums []int) int {
	maxSub := nums[0]
	curSum := 0

	for _, n := range nums {
		if curSum < 0 {
			curSum = 0
		}
		curSum += n
		maxSub = int(math.Max(float64(maxSub), float64(curSum)))
	}

	return maxSub
}
