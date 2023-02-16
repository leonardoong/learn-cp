package main

import "fmt"

func main() {
	fmt.Println(minCostClimbingStairs([]int{10, 15, 20}))
	fmt.Println(minCostClimbingStairs([]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}))

}

func minCostClimbingStairs(cost []int) int {

	cost = append(cost, 0)

	l := len(cost)
	for i := l - 3; i >= 0; i-- {
		cost[i] = min(cost[i]+cost[i+1], cost[i]+cost[i+2])
	}

	return min(cost[0], cost[1])
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
