package main

import "container/list"

/*
 * @lc app=leetcode.cn id=199 lang=golang
 *
 * [199] 二叉树的右视图
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
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := list.New()
	queue.PushBack(root)
	queue.PushBack(nil)
	ans := make([]int, 0)
	i := 0
	for {
		if queue.Front().Value == nil {
			if queue.Front().Next() == nil {
				break
			}
			queue.Remove(queue.Front())
			queue.PushBack(nil)
			i++
			continue
		}
		front := queue.Front()
		if front.Value != nil && front.Next().Value == nil {
			ans = append(ans, front.Value.(*TreeNode).Val)
		}
		if front.Value.(*TreeNode).Left != nil {
			queue.PushBack(front.Value.(*TreeNode).Left)
		}
		if front.Value.(*TreeNode).Right != nil {
			queue.PushBack(front.Value.(*TreeNode).Right)
		}
		queue.Remove(front)
	}
	return ans
}

// @lc code=end
