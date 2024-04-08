package main

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lena, lenb := 0, 0
	ha, hb := headA, headB
	for ha != nil {
		lena++
		ha = ha.Next
	}
	for hb != nil {
		lenb++
		hb = hb.Next
	}
	ha = headA
	hb = headB
	if lena > lenb {
		for i := 0; i < lena-lenb; i++ {
			ha = ha.Next
		}
	} else {
		for i := 0; i < lenb-lena; i++ {
			hb = hb.Next
		}
	}
	for ha != nil && hb != nil {
		if ha == hb {
			return ha
		}
		ha = ha.Next
		hb = hb.Next
	}
	return nil
}

// @lc code=end
