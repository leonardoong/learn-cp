package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%+v \n", spiralOrder([][]int{
		[]int{
			1, 2, 3,
		},
		[]int{
			4, 5, 6,
		},
		[]int{
			7, 8, 9,
		},
	}))

	fmt.Printf("%+v \n", spiralOrder([][]int{
		[]int{
			1, 2, 3, 4,
		},
		[]int{
			5, 6, 7, 8,
		},
		[]int{
			9, 10, 11, 12,
		},
	}))

	fmt.Printf("%+v \n", spiralOrder([][]int{
		[]int{
			1, 2, 3, 4,
		},
		[]int{
			5, 6, 7, 8,
		},
		[]int{
			9, 10, 11, 12,
		},
		[]int{
			13, 14, 15, 16,
		},
	}))
}

func spiralOrder(matrix [][]int) []int {
	res := []int{}
	l := 0
	r := len(matrix[0])
	t := 0
	b := len(matrix)

	for l < r && t < b {

		// add top array (go right)
		for i := l; i < r; i++ {
			res = append(res, matrix[t][i])
		}
		t++

		// go bottom
		for i := t; i < b; i++ {
			res = append(res, matrix[i][r-1])
		}
		r--

		if !(t < b && l < r) {
			break
		}

		// go left
		for i := r - 1; i >= l; i-- {
			res = append(res, matrix[b-1][i])
		}
		b--

		// go up
		for i := b - 1; i >= t; i-- {
			res = append(res, matrix[i][l])
		}
		l++
	}

	return res

}
