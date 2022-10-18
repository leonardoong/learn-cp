package main

import "fmt"

func main() {
	fmt.Println(houseRobber([]int{1, 2, 3, 1}))
	fmt.Println(houseRobber([]int{2, 7, 9, 3, 1}))
}

func houseRobber(nums []int) int {
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
