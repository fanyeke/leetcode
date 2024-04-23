package main

/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (right-left)/2 + left
		num := nums[mid]
		if target == num {
			return mid
		} else if target < num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

// @lc code=end
