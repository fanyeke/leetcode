package main

/*
 * @lc app=leetcode.cn id=128 lang=golang
 *
 * [128] 最长连续序列
 */

// @lc code=start
func longestConsecutive(nums []int) int {
	set := make(map[int]bool)
	for _, v := range nums {
		set[v] = true
	}
	var mmax = 0
	for i := range set {
		if !set[i] {
			continue
		}
		// 不是序列的第一个
		if _, ok := set[i-1]; ok {
			continue
		}
		ans := 1
		for set[i+1] {
			set[i+1] = false
			ans++
			i++
		}
		if ans > mmax {
			mmax = ans
		}
	}
	return mmax
}

// @lc code=end
