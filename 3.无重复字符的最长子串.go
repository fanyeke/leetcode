package main

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var start, ans int
	mmap := make(map[rune]int)
	for i, v := range s {
		if last, found := mmap[v]; found && last >= start {
			start = last + 1
		}
		mmap[v] = i
		ans = max(i-start+1, ans)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end
