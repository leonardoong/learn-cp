package main

import (
	"fmt"
)

func main() {
	prices := []int{2, 4, 1}
	fmt.Println(maxProfit(prices))

	prices = []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(prices))

	prices = []int{7, 6, 4, 3, 1}
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {

	l, r := 0, 1
	maxProf := 0

	for r < len(prices) {
		if prices[l] < prices[r] {
			prof := prices[r] - prices[l]
			if prof > maxProf {
				maxProf = prof
			}
		} else {
			l = r
		}
		r++
	}

	return maxProf
}
