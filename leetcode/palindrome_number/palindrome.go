package main

import "fmt"

func main() {

	x := 121
	fmt.Println(isPalindrome(x))

	x2 := 123
	fmt.Println(isPalindrome(x2))

	x3 := -1232132
	fmt.Println(isPalindrome(x3))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	ori := x
	rev := 0

	for ori != 0 {
		rev = rev*10 + ori%10
		ori = ori / 10
	}

	return x == rev
}
