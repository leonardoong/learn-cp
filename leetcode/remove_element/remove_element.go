package main

import "fmt"

func main() {
	x := []int{1, 1, 2}
	fmt.Println(removeElement(x, 1))

	x2 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeElement(x2, 3))
}

func removeElement(nums []int, val int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	var j = 0

	for i, n := range nums {
		if n != val {
			nums[j] = nums[i]
			j++
		}
	}

	return j
}
