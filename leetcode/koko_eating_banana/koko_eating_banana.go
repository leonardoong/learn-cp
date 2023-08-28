package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(minEatingSpeed([]int{3, 6, 7, 11}, 8))
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 5))
	fmt.Println(minEatingSpeed([]int{30, 11, 23, 4, 20}, 6))
}

func minEatingSpeed(piles []int, h int) int {
	max := 0
	for _, p := range piles {
		if p > max {
			max = p
		}
	}

	l, r := 1, max
	res := r
	for l <= r {
		k := (l + r) / 2
		hours := 0
		for _, p := range piles {
			hours += int(math.Ceil(float64(p) / float64(k)))
		}

		if hours <= h {
			res = min(res, k)
			r = k - 1
		} else {
			l = k + 1
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
