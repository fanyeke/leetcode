package main

/*
 * @lc app=leetcode.cn id=118 lang=golang
 *
 * [118] 杨辉三角
 */

// @lc code=start
func generate(numRows int) [][]int {
	if numRows == 0 {
		return [][]int{}
	}
	ans := make([][]int, 0)
	ans = append(ans, []int{1})
	if numRows == 1 {
		return ans
	}
	ans = append(ans, []int{1, 1})
	if numRows == 2 {
		return ans
	}

	for i := 2; i < numRows; i++ {
		ans = append(ans, make([]int, i+1))
		ans[i][0] = 1
		ans[i][i] = 1
		for j := 1; j < i; j++ {
			ans[i][j] = ans[i-1][j-1] + ans[i-1][j]
		}
	}
	return ans
}

// @lc code=end
