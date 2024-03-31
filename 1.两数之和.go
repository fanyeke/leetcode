package main

import "fmt"

/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	mmap := make(map[int]int, 0)
	for i, v := range nums {
		if ii, ok := mmap[target-v]; ok {
			fmt.Println(mmap)
			fmt.Printf("i = %d; ii = %d", i, ii)
			return []int{min(ii, i), max(ii, i)}
		}
		mmap[v] = i
	}
	return nil
}

// @lc code=end
