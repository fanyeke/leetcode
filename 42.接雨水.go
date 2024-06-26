package main

/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}
	n := len(height)
	res := 0
	// 数组充当备忘录
	l_max := make([]int, n)
	r_max := make([]int, n)
	// 初始化 base case
	l_max[0] = height[0]
	r_max[n-1] = height[n-1]
	// 从左到右计算 l_max
	// 其实也是一种动态规划
	for i := 1; i < n; i++ {
		l_max[i] = max(height[i], l_max[i-1])
	}
	// 从右往左计算 r_max
	for i := n - 2; i >= 0; i-- {
		r_max[i] = max(height[i], r_max[i+1])
	}
	// 计算答案
	for i := 1; i < n-1; i++ {
		res += min(l_max[i], r_max[i]) - height[i]
	}
	return res
}

// @lc code=end
