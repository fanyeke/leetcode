package main

import (
	"fmt"
	"math"
)

/*
 * @lc app=leetcode.cn id=322 lang=golang
 *
 * [322] 零钱兑换
 */

// @lc code=start

func coinChange(coins []int, amount int) int {
	n := len(coins)
	dp := make([]int, amount+1)
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0

	for i := 0; i < n; i++ {
		for j := coins[i]; j <= amount; j++ {
			if dp[j-coins[i]] != math.MaxInt32 {
				dp[j] = min(dp[j], dp[j-coins[i]]+1)
			}
		}
	}
	fmt.Println(dp)
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

// @lc code=end
