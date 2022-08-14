package main

import "fmt"

func main() {
	fmt.Println(isAnagram("rat", "cat"))
	fmt.Println(isAnagram("anagram", "nagaram"))
}

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	countS := make(map[byte]int)
	countT := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		countS[s[i]] = countS[s[i]] + 1
		countT[t[i]] = countT[t[i]] + 1
	}

	for key, val := range countS {
		if val != countT[key] {
			return false
		}
	}

	return true
}
