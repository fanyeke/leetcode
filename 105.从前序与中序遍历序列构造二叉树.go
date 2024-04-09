package main

/*
 * @lc app=leetcode.cn id=105 lang=golang
 *
 * [105] 从前序与中序遍历序列构造二叉树
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
	// 存储 inorder 中值到索引的映射
	valToIndex := make(map[int]int)
	for i, v := range inorder {
		valToIndex[v] = i
	}
	return build(preorder, 0, len(preorder)-1,
		inorder, 0, len(inorder)-1, valToIndex)
}

/*
   定义：前序遍历数组为 preorder[preStart..preEnd]，
   中序遍历数组为 inorder[inStart..inEnd]，
   构造这个二叉树并返回该二叉树的根节点
*/
func build(preorder []int, preStart int, preEnd int,
	inorder []int, inStart int, inEnd int,
	valToIndex map[int]int) *TreeNode {
	if preStart > preEnd {
		return nil
	}

	// root 节点对应的值就是前序遍历数组的第一个元素
	rootVal := preorder[preStart]
	// rootVal 在中序遍历数组中的索引
	index := valToIndex[rootVal]

	leftSize := index - inStart

	// 先构造出当前根节点
	root := &TreeNode{Val: rootVal}
	// 递归构造左右子树
	root.Left = build(preorder, preStart+1, preStart+leftSize,
		inorder, inStart, index-1, valToIndex)

	root.Right = build(preorder, preStart+leftSize+1, preEnd,
		inorder, index+1, inEnd, valToIndex)
	return root
}

// @lc code=end
