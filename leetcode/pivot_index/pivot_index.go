package main

import "fmt"

func main() {
	fmt.Println(pivotIndex([]int{1, 7, 3, 6, 5, 6}))
	fmt.Println(pivotIndex([]int{1, 2, 3}))
	fmt.Println(pivotIndex([]int{2, 1, -1}))
}

func pivotIndex(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}

	count := 0
	for i := 0; i < len(nums); i++ {
		match := total - nums[i] - count
		if match == count {
			return i
		}
		count += nums[i]
	}

	return -1
}
