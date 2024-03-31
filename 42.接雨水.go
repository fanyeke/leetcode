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

	l_max := make([]int, n)
	r_max := make([]int, n)

	l_max[0] = height[0]
	r_max[n-1] = height[n-1]

	for i := 1; i < n; i++ {
		l_max[i] = max(height[i], l_max[i-1])
	}

	for i := n - 2; i >= 0; i-- {
		r_max[i] = max(height[i], r_max[i+1])
	}

	for i := 1; i < n-1; i++ {
		res += min(l_max[i], r_max[i]-height[i])
	}
	return res
}

// @lc code=end
