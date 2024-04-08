package main

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {

	r := head
	l, m := head, head
	i := 1
	for r != nil {
		r = r.Next
		if i > n {
			m = m.Next
		}
		if i > n+1 {
			l = l.Next
		}
		i++
	}
	if l == m {
		head = head.Next
	} else {
		l.Next = m.Next

	}
	return head
}

// @lc code=end
