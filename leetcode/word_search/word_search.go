package main

import (
	"fmt"
)

func main() {
	fmt.Println(exist([][]byte{
		[]byte{
			'A', 'B', 'C', 'E',
		},
		[]byte{
			'S', 'F', 'C', 'S',
		},
		[]byte{
			'A', 'D', 'E', 'E',
		},
	}, "ABCCED"))

	fmt.Println(exist([][]byte{
		[]byte{
			'A', 'B', 'C', 'E',
		},
		[]byte{
			'S', 'F', 'C', 'S',
		},
		[]byte{
			'A', 'D', 'E', 'E',
		},
	}, "SEE"))

	fmt.Println(exist([][]byte{
		[]byte{
			'A', 'B', 'C', 'E',
		},
		[]byte{
			'S', 'F', 'C', 'S',
		},
		[]byte{
			'A', 'D', 'E', 'E',
		},
	}, "ABCB"))
}

func exist(board [][]byte, word string) bool {

	lenRow := len(board)
	lenCol := len(board[0])

	for r := 0; r < lenRow; r++ {
		for c := 0; c < lenCol; c++ {
			if dfs(r, c, 0, board, word) {
				return true
			}
		}
	}

	return false
}

var path = make(map[string]bool)

func dfs(r, c, i int, board [][]byte, word string) bool {
	lenRow := len(board)
	lenCol := len(board[0])

	if i == len(word) {
		return true
	}
	key := fmt.Sprintf("%d+%d", r, c)
	if _, exist := path[key]; exist || r < 0 || c < 0 || r >= lenRow || c >= lenCol || word[i] != board[r][c] {
		return false
	}

	path[key] = true
	res := dfs(r+1, c, i+1, board, word) ||
		dfs(r-1, c, i+1, board, word) ||
		dfs(r, c+1, i+1, board, word) ||
		dfs(r, c-1, i+1, board, word)
	delete(path, key)

	return res
}
