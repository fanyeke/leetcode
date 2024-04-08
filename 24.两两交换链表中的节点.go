package main

/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	v := &ListNode{
		Val:  -1,
		Next: head,
	}
	l, m, r := v, head, head.Next
	head = r
	for {
		m.Next = r.Next
		r.Next = m
		l.Next = r
		if m.Next == nil {
			return head
		}
		m = m.Next
		if r.Next.Next.Next == nil {
			return head
		}
		r = r.Next.Next.Next
		l = l.Next.Next
	}
}

// @lc code=end
