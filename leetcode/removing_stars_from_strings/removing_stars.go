package main

import "fmt"

func main() {
	fmt.Println(removeStars("leet**cod*e"))
	fmt.Println(removeStars("erase*****"))
}

type Stack []string

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() (string, bool) {
	if len(*s) == 0 {
		return "", false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, false
	}
}

func (s *Stack) GetAll() string {
	res := ""
	for _, val := range *s {
		res += val
	}
	return res
}

func removeStars(s string) string {
	stack := Stack{}

	for _, val := range s {
		if val == '*' {
			stack.Pop()
		} else {
			stack.Push(string(val))
		}
	}

	return stack.GetAll()
}
