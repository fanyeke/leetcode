package main

/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func kthSmallest(root *TreeNode, k int) int {
	var dfs func(node *TreeNode)
	i := 0
	ans := 0
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		i++
		if i == k {
			ans = node.Val
		}
		dfs(node.Right)
	}
	dfs(root)
	return ans
}

// @lc code=end
