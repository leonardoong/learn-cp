package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {

	// 1. sort the array
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	res := [][]int{}

	// 2. loop through the array
	for i, n := range nums {
		// 3. if the current number is the same as the previous number, skip it
		if i > 0 && n == nums[i-1] {
			continue
		}

		// 4. find the target number
		l, r := i+1, len(nums)-1
		// 5. use two sums with 2 pointers to find the target number
		for l < r {
			threeSum := n + nums[l] + nums[r]
			if threeSum > 0 {
				r--
			} else if threeSum < 0 {
				l++
			} else {
				res = append(res, []int{n, nums[l], nums[r]})
				// 6. skip the duplicates
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
