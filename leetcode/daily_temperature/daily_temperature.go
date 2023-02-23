package main

import "fmt"

func main() {
	fmt.Printf("res 1 : %+v \n", dailyTemperaturesStack([]int{73, 74, 75, 71, 69, 72, 76, 73}))
	fmt.Printf("res 2 : %+v \n", dailyTemperaturesStack([]int{30, 40, 50, 60}))
	fmt.Printf("res 3 : %+v \n", dailyTemperaturesStack([]int{30, 40, 90}))

}

// brute force
// time  = O(n^2)
func dailyTemperatures(temp []int) []int {
	res := []int{}

	for i := 0; i < len(temp); i++ {
		count := 1
		for j := i + 1; j < len(temp); j++ {
			if temp[i] < temp[j] {
				res = append(res, count)
				break
			}
			count++

			if j == len(temp)-1 {
				res = append(res, 0)
			}
		}

		if i == len(temp)-1 {
			res = append(res, 0)
		}
	}

	return res
}

// use stack
// time : O(n)
func dailyTemperaturesStack(temp []int) []int {
	res := make([]int, len(temp))
	stack := [][]int{}

	for i, v := range temp {

		for len(stack) > 0 && v > stack[len(stack)-1][1] {
			lastIndex := stack[len(stack)-1][0]
			res[lastIndex] = i - lastIndex
			// pop from stack
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, []int{i, v})

	}

	return res
}
