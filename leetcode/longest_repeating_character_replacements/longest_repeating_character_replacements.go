package main

import "fmt"

func main() {
	fmt.Println(characterReplacement("AABABBA", 1))
	fmt.Println(characterReplacement("ABAB", 2))

}

func characterReplacement(s string, k int) int {
	m := make(map[string]int)

	res, l := 0, 0
	h := 0

	for r, char := range s {
		sub := (r - l) + 1

		m[string(char)]++

		if h < m[string(char)] {
			h = m[string(char)]
		}

		if (sub - h) > k {
			m[string(s[l])]--
			l += 1
			continue
		}

		if res < sub {
			res = sub
		}

	}

	return res
}
