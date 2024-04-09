package main

import (
	"container/list"
	"fmt"
)

/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue := list.New()
	queue.PushBack(root)
	queue.PushBack(nil)
	ans := make([][]int, 0)
	ans = append(ans, []int{})
	n := 0
	for queue != nil {
		if queue.Front().Value == nil {
			fmt.Println(ans)
			if queue.Front().Next() == nil {
				break
			}
			ans = append(ans, []int{})
			n++
			queue.Remove(queue.Front())
			queue.PushBack(nil)
		}
		node := queue.Front().Value.(*TreeNode)
		ans[n] = append(ans[n], node.Val)
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
		if node.Right != nil {
			queue.PushBack(node.Right)
		}
		queue.Remove(queue.Front())

	}
	return ans
}

// @lc code=end
