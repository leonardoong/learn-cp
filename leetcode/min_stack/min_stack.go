package main

import "fmt"

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	param_1 := obj.GetMin()
	obj.Pop()
	param_2 := obj.Top()
	param_3 := obj.GetMin()

	fmt.Printf("param 1 = %v, param 2 = %v, param 3 = %v", param_1, param_2, param_3)
}

type MinStack struct {
	val []int
	min []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	n := len(this.val)
	if n == 0 {
		this.val = append(this.val, val)
		this.min = append(this.min, val)
	} else {
		this.val = append(this.val, val)
		minVal := 0
		if len(this.val) == 0 {
			minVal = val
		} else {
			minVal = min(val, this.min[len(this.min)-1])
		}
		this.min = append(this.min, minVal)
	}
}

func (this *MinStack) Pop() {
	this.val = this.val[:len(this.val)-1]
	this.min = this.min[:len(this.min)-1]
}

func (this *MinStack) Top() int {
	return this.val[len(this.val)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
