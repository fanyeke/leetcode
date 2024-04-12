package main

import "fmt"

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start

func backtrack(allPermutations *[][]int, currentPermutation []int, start int, length int) {
	if start == length {
		// 遍历完毕, 拼接答案
		permCopy := make([]int, length)
		copy(permCopy, currentPermutation)
		*allPermutations = append(*allPermutations, permCopy)
		return
	}
	for i := start; i < length; i++ {
		// 交换
		currentPermutation[i], currentPermutation[start] = currentPermutation[start], currentPermutation[i]
		fmt.Println(allPermutations)
		backtrack(allPermutations, currentPermutation, start+1, length)
		// 换回来
		currentPermutation[i], currentPermutation[start] = currentPermutation[start], currentPermutation[i]
	}
}

func permute(nums []int) [][]int {
	var allPermutations [][]int
	backtrack(&allPermutations, nums, 0, len(nums))
	return allPermutations
}

// @lc code=end
