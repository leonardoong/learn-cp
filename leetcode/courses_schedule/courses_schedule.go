package main

import (
	"fmt"
)

func main() {
	fmt.Println(canFinish(2, [][]int{
		[]int{
			1, 0,
		},
	}))

	fmt.Println(canFinish(2, [][]int{
		[]int{
			1, 0,
		},
		[]int{
			0, 1,
		},
	}))

	fmt.Println(canFinish(2, [][]int{
		[]int{
			0, 1,
		},
		[]int{
			0, 2,
		},
		[]int{
			1, 3,
		},
		[]int{
			3, 4,
		},
		[]int{
			1, 4,
		},
	}))
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	var visited = make(map[int]bool)
	var preMap = make(map[int][]int)
	fmt.Printf("preMap === %v \n", preMap)

	for i := 0; i < numCourses; i++ {
		preMap[i] = nil
	}

	for _, req := range prerequisites {
		preMap[req[0]] = append(preMap[req[0]], req[1])
	}

	for i := 0; i < numCourses; i++ {
		if !dfs(i, preMap, visited) {
			return false
		}
	}

	return true
}

func dfs(crs int, preMap map[int][]int, visited map[int]bool) bool {
	if visited[crs] {
		return false
	}

	if preMap[crs] == nil {
		return true
	}

	visited[crs] = true
	for _, pre := range preMap[crs] {
		if !dfs(pre, preMap, visited) {
			return false
		}
	}

	delete(visited, crs)
	preMap[crs] = nil
	return true
}
