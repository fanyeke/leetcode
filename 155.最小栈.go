package main

/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
type MinStack struct {
	stack    []int
	minstack []int
}

func Constructor() MinStack {
	return MinStack{
		stack:    make([]int, 0),
		minstack: make([]int, 0),
	}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if len(this.minstack) == 0 || val <= this.minstack[len(this.minstack)-1] {
		this.minstack = append(this.minstack, val)
	}
}

func (this *MinStack) Pop() {
	if this.stack[len(this.stack)-1] == this.minstack[len(this.minstack)-1] {
		this.minstack = this.minstack[:len(this.minstack)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minstack[len(this.minstack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
