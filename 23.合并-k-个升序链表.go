package main

import "container/heap"

/*
 * @lc app=leetcode.cn id=23 lang=golang
 *
 * [23] 合并 K 个升序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// PriorityQueue 为了使用堆操作定义的优先队列
type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// 按值的升序排列
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*ListNode)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

// mergeKLists 使用最小堆合并K个链表
func mergeKLists(lists []*ListNode) *ListNode {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	// 将k个链表的头节点全部加入优先队列
	for _, list := range lists {
		if list != nil {
			heap.Push(&pq, list)
		}
	}

	// dummy头节点和current指针用于构建结果链表
	dummy := &ListNode{}
	current := dummy
	for pq.Len() > 0 {
		// 弹出最小的节点，连接到结果链表
		minNode := heap.Pop(&pq).(*ListNode)
		current.Next = minNode
		current = current.Next

		// 如果最小节点的next不为空，将其加入优先队列
		if minNode.Next != nil {
			heap.Push(&pq, minNode.Next)
		}
	}

	return dummy.Next
}

// @lc code=end
