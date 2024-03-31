package main

/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	r := 0
	for i := 0; i < len(nums); i++ {
		nums[i], nums[r] = nums[r], nums[i]
		if nums[r] != 0 {
			r++
		}
	}
}

// @lc code=end
