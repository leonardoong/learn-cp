package main

import "fmt"

func main() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
	fmt.Println(isSubsequence("axc", "ahbgdc"))
}

func isSubsequence(s string, t string) bool {

	i := 0
	j := 0
	for j < len(t) && i < len(s) {
		if s[i] == t[j] {
			i++
			j++
		} else {
			j++
		}
	}

	if i == len(s) {
		return true
	}

	return false
}
