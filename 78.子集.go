package main

/*
 * @lc app=leetcode.cn id=78 lang=golang
 *
 * [78] 子集
 */

// @lc code=start
func subsets(nums []int) [][]int {
	res := [][]int{}
	track := []int{}
	btrack(nums, 0, track, &res)
	return res
}

func btrack(nums []int, start int, track []int, res *[][]int) {
	temp := make([]int, len(track))
	copy(temp, track)
	*res = append(*res, temp)
	for i := start; i < len(nums); i++ {
		// 做选择
		track = append(track, nums[i])
		// 回溯
		btrack(nums, i+1, track, res)
		// 撤销选择
		track = track[:len(track)-1]
	}
}

// @lc code=end
