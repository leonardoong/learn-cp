package main

import (
	"fmt"
	"strings"
)

func main() {

	x := "III"
	fmt.Println(romanToInt(x))

	x2 := "LVIII"
	fmt.Println(romanToInt(x2))

	x3 := "MCMXCIV"
	fmt.Println(romanToInt(x3))
}

func romanToInt(x string) int {
	romanMap := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	var result int
	var used string
	for i, s := range x {
		l := string(s)

		if l == "I" || l == "X" || l == "C" {
			next := i + 1
			if i == len(x)-1 {
				next = i
			}

			c := l + string(x[next])
			_, ok := romanMap[c]
			if ok {
				result += romanMap[c]
				used += c
			} else {
				if !strings.Contains(used, l) {
					result += romanMap[l]
				}
			}
		} else {
			if !strings.Contains(used, l) {
				result += romanMap[l]
			}
		}
	}

	return result
}
