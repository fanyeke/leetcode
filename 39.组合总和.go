package main

/*
 * @lc app=leetcode.cn id=39 lang=golang
 *
 * [39] 组合总和
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	var res [][]int

	// 定义回溯函数
	var backtrack func(start int, target int, sum int, track []int)

	backtrack = func(start int, target int, sum int, track []int) {
		// 如果当前 sum 等于 target，说明已找到符合要求的组合
		if sum == target {
			// 由于 track 数组为引用类型，因此需要重新生成一个数组对象来保存到结果集
			res = append(res, append([]int{}, track...))
			return
		}

		// 如果当前 sum 已经大于 target，回溯结束
		if sum > target {
			return
		}

		// 从指定位置开始遍历候选数组
		for i := start; i < len(candidates); i++ {
			// 将候选数组当前元素加入路径
			track = append(track, candidates[i])
			sum += candidates[i]
			// 继续遍历下一个元素
			backtrack(i, target, sum, track)
			// 回溯：将当前元素从路径中移除
			sum -= candidates[i]
			track = track[:len(track)-1]
		}
	}

	backtrack(0, target, 0, []int{})

	return res
}

// @lc code=end
