package main

import "fmt"

func main() {
	fmt.Println(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 1))
	fmt.Println(topKFrequent([]int{1, 2}, 2))
	fmt.Println(topKFrequent([]int{1}, 1))
}

func topKFrequent(nums []int, k int) []int {
	countMap := make(map[int]int)
	freq := make([][]int, len(nums)+1)

	// 1. Count the frequency of each number
	for _, num := range nums {
		countMap[num] = 1 + countMap[num]
	}

	// 2. Store the frequency in a 2D array, key is the frequency, value is the number
	// eg. freq[3] = []int{1}
	//     freq[2] = []int{2}
	for num, count := range countMap {
		freq[count] = append(freq[count], num)
	}

	res := []int{}
	// 3. Loop backwards from the highest frequency to the lowest frequency
	for i := len(freq) - 1; i > 0; i-- {
		for _, n := range freq[i] {
			// 4. Add the number to the result if the length of the result is less than k
			res = append(res, n)
			if len(res) == k {
				return res
			}
		}
	}

	return res
}
