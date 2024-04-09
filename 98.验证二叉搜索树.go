package main

/*
 * @lc app=leetcode.cn id=98 lang=golang
 *
 * [98] 验证二叉搜索树
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
func isValidBST(root *TreeNode) bool {
	return isValid(root, nil, nil)
}

/* 限定以 root 为根的子树节点必须满足 max.val > root.val > min.val */
func isValid(root *TreeNode, min *TreeNode, max *TreeNode) bool {
	// base case
	if root == nil {
		return true
	}
	// 若 root.val 不符合 max 和 min 的限制，说明不是合法 BST
	if min != nil && root.Val <= min.Val {
		return false
	}
	if max != nil && root.Val >= max.Val {
		return false
	}
	// 限定左子树的最大值是 root.Val，右子树的最小值是 root.Val
	return isValid(root.Left, min, root) && isValid(root.Right, root, max)
}

// @lc code=end
