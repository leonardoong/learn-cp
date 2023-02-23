package main

import "fmt"

func main() {
	fmt.Println(longestConsecutive([]int{100, 4, 200, 1, 3, 2}))
	fmt.Println(longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}))

}

func longestConsecutive(nums []int) int {
	numMap := make(map[int]bool)
	longest := 0

	for _, n := range nums {
		numMap[n] = true
	}

	for _, n := range nums {
		if _, exist := numMap[n-1]; !exist {
			length := 0
			for {
				if _, exist2 := numMap[n+length]; exist2 {
					length++
				} else {
					break
				}
			}
			longest = max(length, longest)

		}
	}

	return longest
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
