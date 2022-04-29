package main

import "fmt"

func main() {
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("(]"))
}

func isValid(str string) bool {
	arr := []string{}
	pair := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}

	for _, s := range str {

		ss := string(s)
		if ss == "(" || ss == "{" || ss == "[" {
			arr = append(arr, ss)
		} else {
			if len(arr) == 0 || arr[len(arr)-1] != pair[ss] {
				return false
			}
			arr = arr[:len(arr)-1]
		}

	}
	return len(arr) == 0
}
