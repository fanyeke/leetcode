package main

import "fmt"

/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	h := head
	length := 0
	for h != nil {
		length++
		h = h.Next
	}
	m := head
	for i := 0; i < length/2; i++ {
		fmt.Println("m=", m.Val)
		m = m.Next
	}
	var r, l *ListNode
	l, r = nil, m.Next
	for r != nil {
		m.Next = l
		l = m
		m = r
		r = r.Next
	}
	m.Next = l
	// m1 := m
	// for m1 != nil {
	// 	fmt.Printf("%d\t", m1.Val)
	// }
	h = head
	for i := 0; i < length/2; i++ {
		fmt.Printf("m.Val=%d, h.Val=%d\n", m.Val, h.Val)
		if m.Val != h.Val {
			return false
		}
		m = m.Next
		h = h.Next
	}
	return true
}

// @lc code=end
