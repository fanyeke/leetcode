package main

import "fmt"

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

/*
 * @lc app=leetcode.cn id=138 lang=golang
 *
 * [138] 随机链表的复制
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}

	h := head
	nodeSlice := make([]int, 0)
	randomSlice := make([]*Node, 0)
	for h != nil {
		randomSlice = append(randomSlice, h)
		h = h.Next
	}
	for i := 0; i < len(randomSlice); i++ {
		if randomSlice[i].Random == nil {
			nodeSlice = append(nodeSlice, -1)
			continue
		}
		for j := 0; j < len(randomSlice); j++ {
			if randomSlice[i].Random == randomSlice[j] {
				nodeSlice = append(nodeSlice, j)
				break
			}
		}
	}
	newHead := &Node{
		Val:  head.Val,
		Next: nil,
	}
	nh := newHead
	for i := 1; i < len(randomSlice); i++ {
		nh.Next = &Node{
			Val:  randomSlice[i].Val,
			Next: nil,
		}
		nh = nh.Next
	}
	fmt.Println(randomSlice)
	fmt.Println(nodeSlice)
	nh = newHead
	for i := 0; i < len(nodeSlice); i++ {
		if nodeSlice[i] == -1 {
			nh = nh.Next
			continue
		}
		rh := newHead
		for j := 1; j <= nodeSlice[i]; j++ {
			rh = rh.Next
		}
		nh.Random = rh
		nh = nh.Next
	}
	return newHead
}

// @lc code=end
