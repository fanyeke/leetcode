package main

import "math"

/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
func maxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	dp := make([]int, n)

	dp[0] = nums[0]

	for i := 1; i < n; i++ {
		dp[i] = max(nums[i], nums[i]+dp[i-1])
	}
	res := math.MinInt32
	for i := 0; i < n; i++ {
		res = max(res, dp[i])
	}
	return res
}

// @lc code=end
