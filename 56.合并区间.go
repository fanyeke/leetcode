package main

import (
	"sort"
)

/*
 * @lc app=leetcode.cn id=56 lang=golang
 *
 * [56] 合并区间
 */

// @lc code=start
func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return [][]int{}
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans := make([][]int, 0)
	mmax := intervals[0][1]
	mmin := intervals[0][0]
	for i := 1; i < len(intervals); i++ {
		if mmax >= intervals[i][0] {
			mmax = max(mmax, intervals[i][1])
		} else {
			ans = append(ans, []int{mmin, mmax})
			mmin = intervals[i][0]
			mmax = intervals[i][1]
		}
	}
	ans = append(ans, []int{mmin, mmax})
	return ans
}

// @lc code=end
