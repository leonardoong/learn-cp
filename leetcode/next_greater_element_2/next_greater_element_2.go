package main

import "fmt"

func main() {

	fmt.Println(nextGreaterElements([]int{1, 2, 1}))
	fmt.Println(nextGreaterElements([]int{1, 2, 3, 4, 3}))

	fmt.Println(nextGreaterElements2([]int{1, 2, 1}))
	fmt.Println(nextGreaterElements2([]int{1, 2, 3, 4, 3}))

}

// brute force 1
// TC : O(n^2)
// SC : O(n)
func nextGreaterElements(nums []int) []int {
	result := make([]int, len(nums))

	doubleSlice := []int{}
	doubleSlice = append(doubleSlice, nums...)
	doubleSlice = append(doubleSlice, nums...)

	for i := 0; i < len(nums); i++ {
		result[i] = -1
		for j := i + 1; j < len(doubleSlice); j++ {
			if doubleSlice[i] < doubleSlice[j] {
				result[i] = doubleSlice[j]
				break
			}
		}
	}

	return result
}

// brute force 2 with modulus
// TC : O(n^2)
// SC : O(n)
func nextGreaterElements2(nums []int) []int {
	result := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		result[i] = -1
		for j := 1; j < len(nums); j++ {
			if nums[(i+j)%len(nums)] > nums[i] {
				result[i] = nums[(i+j)%len(nums)]
				break
			}
		}
	}

	return result
}
