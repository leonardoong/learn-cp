package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{9, 0, 3, 5, 7}))
}

func subsets(nums []int) [][]int {
	var queue [][]int
	queue = append(queue, []int{})

	for i := 0; i < len(nums); i++ {
		length := len(queue)

		for j := 0; j < length; j++ {
			item := queue[j]

			tmp := make([]int, len(item))
			copy(tmp, item)
			item = append(item, nums[i])
			queue = append(queue, item)
		}
	}

	return queue
}
