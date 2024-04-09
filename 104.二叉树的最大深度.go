package main

import "math"

/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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
func maxDepth(root *TreeNode) int {
	ans := math.MinInt

	var dfs func(node *TreeNode, depth int)
	dfs = func(node *TreeNode, depth int) {
		ans = max(ans, depth)
		if node == nil {
			return
		}
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)

	}
	dfs(root, 0)
	return ans
}

// @lc code=end
