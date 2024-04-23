package main

import "fmt"

/*
 * @lc app=leetcode.cn id=33 lang=golang
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		num := nums[mid]
		fmt.Println(mid)

		if num > nums[left] {
			left = mid + 1
		} else if num < nums[left] {
			right = mid - 1
		} else {
			return left
		}
	}
	return left
}

// @lc code=end
