package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	var queue [][]int
	queue = append(queue, []int{})

	for i := 0; i < len(nums); i++ {
		length := len(queue)
		fmt.Println(length)

		for j := 0; j < length; j++ {
			item := queue[j]

			item = append(item, nums[i])
			queue = append(queue, item)
		}
	}

	return queue
}
