package main

import (
	"fmt"
)

func main() {
	fmt.Println(searchMatrix1([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}, 3))

	fmt.Println(searchMatrix2([][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}, 3))
}

func searchMatrix1(matrix [][]int, target int) bool {

	res := false

	if len(matrix) == 1 {
		for _, v := range matrix[0] {
			if v == target {
				res = true
				return res
			}
		}
		return res
	}

	for i, row := range matrix {
		// find closest row  to target or last row
		if row[0] > target {
			indexToLoop := max(i-1, 0)
			for _, v := range matrix[indexToLoop] {
				if v == target {
					res = true
					return res
				}
			}
		}

		// loop for last row, final check
		if i == len(matrix)-1 {
			for _, v := range matrix[i] {
				if v == target {
					res = true
					return res
				}
			}
		}
	}

	return res

}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func searchMatrix2(matrix [][]int, target int) bool {
	left, right := 0, len(matrix)*len(matrix[0])
	for left < right {

		middle := (left + right) / 2

		i, j := middle/len(matrix[0]), middle%len(matrix[0])

		if target < matrix[i][j] {
			right = middle
			continue
		}

		if target > matrix[i][j] {
			left = middle + 1
			continue
		}

		return true
	}

	return false
}
