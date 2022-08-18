package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	res := [][]int{}

	for i, n := range nums {
		if i > 0 && n == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			threeSum := n + nums[l] + nums[r]
			if threeSum > 0 {
				r--
			} else if threeSum < 0 {
				l++
			} else {
				res = append(res, []int{n, nums[l], nums[r]})
				l++
				for nums[l] == nums[l-1] {
					if l < r {
						l++
					}
				}
			}
		}
	}

	return res
}
