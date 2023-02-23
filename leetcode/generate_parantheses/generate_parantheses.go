package main

import "fmt"

func main() {
	fmt.Printf("generate : %+v \n", generateParenthesis(3))
}

func backtrack(res *[]string, n, open, close int, curSequence string) {
	if open == close && close == n {
		*res = append(*res, curSequence)
		return
	}

	if open < n {
		curSequence += "("
		open++
		backtrack(res, n, open, close, curSequence)
		open--
		curSequence = curSequence[:len(curSequence)-1]
	}

	if close < open {
		curSequence += ")"
		close++
		backtrack(res, n, open, close, curSequence)
		close--
		curSequence = curSequence[:len(curSequence)-1]
	}
}

func generateParenthesis(n int) []string {
	var res []string
	backtrack(&res, n, 0 /*open*/, 0 /*close*/, "" /*curSequence*/)
	return res
}
