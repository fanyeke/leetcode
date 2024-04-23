package main

import "fmt"

/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	left, right := 0, len(matrix)-1
	for left <= right {
		mid := (right-left)/2 + left
		num := matrix[mid][0]
		if target < num {
			right = mid - 1
		} else if target > num {
			left = mid + 1
		} else {
			return true
		}
	}
	left--
	if left < 0 {
		return false
	}
	fmt.Println(left)
	fmt.Println(matrix[left])
	l, r := 0, len(matrix[left])-1
	for l <= r {
		mid := (r-l)/2 + l
		num := matrix[left][mid]
		if target < num {
			r = mid - 1
		} else if target > num {
			l = mid + 1
		} else {
			return true
		}
	}
	return false

}

// @lc code=end
