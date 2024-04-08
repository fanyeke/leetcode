package main

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var head *ListNode
	var h *ListNode
	head = &ListNode{Val: 9999}
	h = head
	for list1 != nil || list2 != nil {
		if list1 == nil {
			for list2 != nil {
				h.Next = &ListNode{Val: list2.Val}
				h = h.Next
				list2 = list2.Next
			}
			break
		}
		if list2 == nil {
			for list1 != nil {
				h.Next = &ListNode{Val: list1.Val}
				h = h.Next
				list1 = list1.Next
			}
			break
		}
		if list1.Val < list2.Val {
			// fmt.Println("11111")
			h.Next = &ListNode{Val: list1.Val}
			h = h.Next
			list1 = list1.Next
		} else {
			// fmt.Println("2222")
			h.Next = &ListNode{Val: list2.Val}
			h = h.Next
			list2 = list2.Next
		}
	}
	h = head
	for h != nil {
		// fmt.Printf("%d\t", h.Val)
		h = h.Next
	}
	return head.Next
}

// @lc code=end
