package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%v\n", groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
	fmt.Printf("%v\n", groupAnagrams([]string{""}))
	fmt.Printf("%v\n", groupAnagrams([]string{
		"bdddddddddd", "bbbbbbbbbbc",
	}))

}

func groupAnagrams(strs []string) [][]string {
	resMap := make(map[string][]string)

	// 1. loop through the strs
	for _, s := range strs {
		count := make([]int, 26)

		// 2. loop through each char in the string
		for _, b := range s {
			// 3. set the count of each char
			asci := int(b)
			ascia := int('a')
			count[asci-ascia] = count[asci-ascia] + 1
		}

		// 4. convert the count to string for key map
		key := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(count)), "#"), "[]")

		// 5. append the string to the key map
		resMap[key] = append(resMap[key], s)
	}

	// 6. convert the map to array
	res := [][]string{}
	for _, mValue := range resMap {
		res = append(res, mValue)
	}

	return res
}
