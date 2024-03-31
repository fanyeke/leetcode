package main

/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为 K 的子数组
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		sum := nums[i]
		if sum == k {
			ans++
		}
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum == k {
				ans++
			}
		}
	}
	return ans
}

// @lc code=end
