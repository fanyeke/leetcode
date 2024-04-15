package main

import "fmt"

/*
 * @lc app=leetcode.cn id=73 lang=golang
 *
 * [73] 矩阵置零
 */

// @lc code=start
func setZeroes(matrix [][]int) {
	x := make(map[int]bool)
	y := make(map[int]bool)
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				x[i] = true
				y[j] = true
			}
		}
	}
	for k, _ := range x {
		for j := 0; j < len(matrix[k]); j++ {
			matrix[k][j] = 0
		}
	}
	for k, _ := range y {
		for i := 0; i < len(matrix); i++ {
			matrix[i][k] = 0
		}
	}
	fmt.Println(matrix)
}

// @lc code=end
