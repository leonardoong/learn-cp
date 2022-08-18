package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 1, 1, 2, 2, 3}))
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	fmt.Println(maxArea([]int{1, 2}))
	fmt.Println(maxArea([]int{1, 2, 1}))
}

func maxArea(height []int) int {
	l, r := 0, len(height)-1
	res := 0
	for l < r {
		minHeight := 0
		if height[l] > height[r] {
			minHeight = height[r]
		} else {
			minHeight = height[l]
		}

		area := (r - l) * minHeight
		if res < area {
			res = area
		}

		if minHeight == height[l] {
			l++
		} else {
			r--
		}
	}

	return res
}
