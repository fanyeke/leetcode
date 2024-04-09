package main

/*
 * @lc app=leetcode.cn id=114 lang=golang
 *
 * [114] 二叉树展开为链表
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
func flatten(root *TreeNode) {
	// var findLeft func(node *TreeNode)
	// sum := 0
	// left := &TreeNode{}
	// findLeft = func(node *TreeNode) {
	// 	sum++
	// 	if node.Left == nil {
	// 		left = node
	// 		return
	// 	}
	// 	findLeft(node.Left)
	// }
	// findLeft(root)
	if root == nil {
		return
	}
	h := root
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		h.Left = node
		h = h.Left
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root.Left)
	dfs(root.Right)
	var ddfs func(node *TreeNode)
	ddfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		ddfs(node.Left)
		node.Right = node.Left
		node.Left = nil
	}
	ddfs(root)
}

// @lc code=end
