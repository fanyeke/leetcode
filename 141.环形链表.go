package main

/*
 * @lc app=leetcode.cn id=141 lang=golang
 *
 * [141] 环形链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for {
		if slow == nil || fast == nil {
			return false
		}
		fast = fast.Next
		if fast == slow {
			return true
		}
		if slow == nil || fast == nil {
			return false
		}
		fast = fast.Next
		if fast == slow {
			return true
		}
		slow = slow.Next
	}
}

// @lc code=end
