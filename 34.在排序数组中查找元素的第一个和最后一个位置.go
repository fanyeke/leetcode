package main

/*
 * @lc app=leetcode.cn id=34 lang=golang
 *
 * [34] 在排序数组中查找元素的第一个和最后一个位置
 */

// @lc code=start
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left, right := 0, len(nums)-1
	if nums[left] > target || nums[right] < target {
		return []int{-1, -1}
	}
	for left <= right {
		mid := (-left+right)/2 + left
		num := nums[mid]
		if target == num {
			left = mid
			break
		} else if target < num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	min, max := left, left
	if nums[left] != target {
		return []int{-1, -1}

	}
	for i := left; i < len(nums); i++ {
		if nums[i] == target {
			max = i
		} else {
			break
		}
	}
	for i := left; i >= 0; i-- {
		if nums[i] == target {
			min = i
		} else {
			break
		}
	}
	return []int{min, max}
}

// @lc code=end
