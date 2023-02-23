package main

import (
	"fmt"
	"sort"
)

func main() {
	// fmt.Println(carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}))
	fmt.Println(carFleet(10, []int{6, 8}, []int{3, 2}))

}

type car struct {
	Position int
	Speed    int
}

// time : O(nlogn)
// memory : O(n)
func carFleet(target int, position []int, speed []int) int {
	cars := []car{}

	for i := 0; i < len(position); i++ {
		cars = append(cars, car{
			Position: position[i],
			Speed:    speed[i],
		})
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].Position < cars[j].Position
	})

	stack := []float64{}
	for i := len(cars) - 1; i >= 0; i-- {
		stack = append(stack, float64((target-cars[i].Position))/float64(cars[i].Speed))

		if len(stack) >= 2 && stack[len(stack)-1] <= stack[len(stack)-2] {
			// pop stack
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack)
}
