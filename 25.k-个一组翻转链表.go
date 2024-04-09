package main

/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	a, b := head, head

	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}
	newHead := reverse(a, b)

	a.Next = reverseKGroup(b, k)
	return newHead
}

func reverse(a, b *ListNode) *ListNode {
	pre, cur, nxt := (*ListNode)(nil), a, a
	for cur != b {
		nxt = cur.Next
		cur.Next = pre
		pre, cur = cur, nxt
	}
	return pre
}

// @lc code=end
