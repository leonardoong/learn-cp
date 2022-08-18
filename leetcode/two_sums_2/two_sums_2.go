package main

import "fmt"

func main() {

	num1 := []int{2, 7, 11, 15}
	target1 := 9
	num2 := []int{2, 3, 4}
	target2 := 6
	num3 := []int{3, 3}
	target3 := 6
	fmt.Println(twoSum(num1, target1))
	fmt.Println(twoSum(num2, target2))
	fmt.Println(twoSum(num3, target3))

}

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	for l < r {
		if numbers[l]+numbers[r] > target {
			r--
		} else if numbers[l]+numbers[r] < target {
			l++
		} else if numbers[l]+numbers[r] == target {
			return []int{l + 1, r + 1}
		}
	}

	return []int{}
}
