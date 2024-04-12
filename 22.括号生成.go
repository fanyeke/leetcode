package main

/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	var dfs func(i, j int, s string)
	dfs = func(i, j int, s string) {
		if j > n || i > n || j > i {
			return
		}
		dfs(i+1, j, s+"(")
		dfs(i, j+1, s+")")
		if i == n && j == n {
			ans = append(ans, s)
		}
	}
	dfs(0, 0, "")
	return ans
}

// @lc code=end
