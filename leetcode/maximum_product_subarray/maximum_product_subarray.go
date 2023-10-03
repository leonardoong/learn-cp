package main

import "fmt"

func main() {
	fmt.Println(maxProduct([]int{2, 3, -2, 4}))
	fmt.Println(maxProduct([]int{-2, 0, -1}))
}

func maxProduct(nums []int) int {
	res := nums[0]
	curMin, curMax := nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		tmp := curMax * nums[i]
		curMax = MaxOf(nums[i]*curMax, nums[i]*curMin, nums[i])
		curMin = MinOf(tmp, nums[i]*curMin, nums[i])
		res = MaxOf(res, curMax)
	}

	return res
}

func MinOf(vars ...int) int {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}

func MaxOf(vars ...int) int {
	max := vars[0]

	for _, i := range vars {
		if max < i {
			max = i
		}
	}

	return max
}
