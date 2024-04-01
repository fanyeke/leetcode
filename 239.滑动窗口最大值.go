package main

/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start

// 单调队列的实现
type MonotonicQueue struct {
	q []int
}

func (mq *MonotonicQueue) push(n int) {
	for len(mq.q) > 0 && mq.q[len(mq.q)-1] < n {
		mq.q = mq.q[:len(mq.q)-1]
	}
	mq.q = append(mq.q, n)
}

func (mq *MonotonicQueue) max() int {
	return mq.q[0]
}

func (mq *MonotonicQueue) pop(n int) {
	if n == mq.q[0] {
		mq.q = mq.q[1:]
	}
}

func maxSlidingWindow(nums []int, k int) []int {
	window := MonotonicQueue{make([]int, 0)}
	res := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		if i < k-1 {
			window.push(nums[i])
		} else {
			window.push(nums[i])
			res = append(res, window.max())
			window.pop(nums[i-k+1])
		}
		// fmt.Println(window.q)
	}
	return res
}

// @lc code=end
