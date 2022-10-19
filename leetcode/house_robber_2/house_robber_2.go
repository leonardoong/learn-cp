package main

import "fmt"

func main() {
	fmt.Println(rob([]int{1, 2, 3, 1}))
	fmt.Println(rob([]int{2, 7, 9, 3, 1}))
	fmt.Println(rob([]int{2, 3, 2}))
}

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	first := rob2(nums[1:])
	second := rob2(nums[:len(nums)-1])

	return max(first, second)
}

func rob2(nums []int) int {
	rob1, rob2 := 0, 0

	temp := 0

	for _, v := range nums {
		temp = max(v+rob1, rob2)
		rob1 = rob2
		rob2 = temp
	}

	return temp
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
