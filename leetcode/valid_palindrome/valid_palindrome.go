package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

func main() {
	fmt.Println(isPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(isPalindrome("race a car"))
}

func isPalindrome(s string) bool {
	// 1. Remove all non-alphanumeric characters
	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	str := reg.ReplaceAllString(s, "")
	// 2. Convert to lower case
	lowerStr := strings.ToLower(str)

	//	3. Check if it is a palindrome
	l, r := 0, len(lowerStr)-1
	for l < r {
		if lowerStr[l] != lowerStr[r] {
			return false
		}
		l++
		r--
	}

	return true
}
