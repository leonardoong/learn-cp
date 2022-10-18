package main

import "fmt"

func main() {
	fmt.Println(climbStairs(5))
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(10))
}

func climbStairs(n int) int {
	// 1D DP
	x, y := 1, 1

	for i := 0; i < n-1; i++ {
		temp := x
		x = x + y
		y = temp
	}
	return x
}
