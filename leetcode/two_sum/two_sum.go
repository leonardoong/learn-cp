package main

import "fmt"

func main() {

	num1 := []int{2, 7, 11, 15}
	target1 := 9
	num2 := []int{3, 2, 4}
	target2 := 6
	num3 := []int{3, 3}
	target3 := 6
	fmt.Println(twoSum(num1, target1))
	fmt.Println(twoSum(num2, target2))
	fmt.Println(twoSum(num3, target3))

}

func twoSum(nums []int, target int) []int {

	seen := make(map[int]int)

	for i, num := range nums {
		if _, ok := seen[target-num]; ok {
			return []int{seen[target-num], i}
		}
		seen[num] = i
	}

	return []int{}
}
