package main

/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	res := 0
	for left < right {
		// 求之间的矩形面积
		curArea := func() int {
			if height[left] < height[right] {
				return height[left] * (right - left)
			}
			return height[right] * (right - left)
		}()
		res = func() int {
			if curArea > res {
				return curArea
			}
			return res
		}()
		// 移动较低的一边
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return res
}

// @lc code=end
