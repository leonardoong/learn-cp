package main

import (
	"fmt"
)

func main() {
	fmt.Println(isIsomorphic("agg", "add"))
	fmt.Println(isIsomorphic("foo", "bar"))
	fmt.Println(isIsomorphic("paper", "title"))

}

func isIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	mapS := make(map[string]string)
	mapT := make(map[string]string)

	for i := 0; i < len(s); i++ {
		strS := string(s[i])
		strT := string(t[i])
		_, ok := mapS[strS]
		_, ok2 := mapT[strT]
		if !ok && !ok2 {
			mapS[strS] = strT
			mapT[strT] = strS
		} else if mapS[strS] != strT && mapT[strT] != strS {
			return false
		}
	}

	return true
}
