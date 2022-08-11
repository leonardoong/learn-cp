package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(nums))

	nums = []int{7, 1, 5, 3, 6, 4}
	fmt.Println(containsDuplicate(nums))

	nums = []int{1, 1, 1, 3, 3, 4, 3, 2, 4}
	fmt.Println(containsDuplicate(nums))
}

func containsDuplicate(nums []int) bool {

	tempMap := map[int]bool{}
	for _, num := range nums {
		if _, exist := tempMap[num]; exist {
			return true
		}
		tempMap[num] = true
	}

	return false
}
