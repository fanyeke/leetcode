package main

import "sort"

/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	h := head
	rs := make([]int, 0)
	for h != nil {
		rs = append(rs, h.Val)
		h = h.Next
	}
	sort.Ints(rs)
	h = head
	i := 0
	for h != nil {
		h.Val = rs[i]
		i++
		h = h.Next
	}
	return head
}

// @lc code=end
