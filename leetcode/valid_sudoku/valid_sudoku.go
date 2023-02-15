package main

import (
	"fmt"
)

func main() {
	fmt.Println(isValidSudoku([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}))

	fmt.Println(isValidSudoku([][]byte{
		{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}))
}

func isValidSudoku(board [][]byte) bool {

	var (
		cols   = make(map[int]map[byte]bool)
		rows   = make(map[int]map[byte]bool)
		square = make(map[string]map[byte]bool)
	)

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}

			squareBox := fmt.Sprintf("%v_%v", r/3, c/3)

			_, existCol := cols[c][board[r][c]]
			_, existRow := rows[r][board[r][c]]
			_, existSquare := square[squareBox][board[r][c]]

			if existCol || existRow || existSquare {
				return false
			}

			if cols[c] == nil {
				cols[c] = make(map[byte]bool)
			}
			if rows[r] == nil {
				rows[r] = make(map[byte]bool)
			}
			if square[squareBox] == nil {
				square[squareBox] = make(map[byte]bool)
			}

			cols[c][board[r][c]] = true
			rows[r][board[r][c]] = true
			square[squareBox][board[r][c]] = true

		}
	}
	return true
}
