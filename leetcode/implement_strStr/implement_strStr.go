package main

import "fmt"

func main() {
	var h, n = "", ""
	h = "hello"
	n = "ll"
	fmt.Println(strStr(h, n))

	h = "aaaaa"
	n = "bba"
	fmt.Println(strStr(h, n))
}

func strStr(haystack string, needle string) int {

	if len(needle) == 0 {
		return 0
	}

	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			if len(needle) == 1 {
				return i
			}
			for j := 1; j < len(needle); j++ {
				if i+j >= len(haystack) || haystack[i+j] != needle[j] {
					break
				}
				if j == len(needle)-1 {
					return i
				}
			}
		}
	}

	return -1
}
