package main

import "fmt"

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
}

func lengthOfLongestSubstring(s string) int {
	m := make(map[rune]int)
	curr, max, l := 0, 0, 0

	for i, char := range s {

		if val, ok := m[char]; ok && val > l {
			curr = i - val
			l = val
		} else if !ok || val < l {
			curr++
		}

		if curr > max {
			max = curr
		}

		m[char] = i

	}

	return max
}
