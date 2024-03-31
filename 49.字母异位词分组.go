package main

/*
 * @lc app=leetcode.cn id=49 lang=golang
 *
 * [49] 字母异位词分组
 */

// @lc code=start
func groupAnagrams(strs []string) [][]string {
	codeToGroup := make(map[string][]string)
	for _, s := range strs {
		code := encode(s)
		codeToGroup[code] = append(codeToGroup[code], s)
	}
	res := make([][]string, 0, len(codeToGroup))
	for _, group := range codeToGroup {
		res = append(res, group)
	}
	return res
}

func encode(s string) string {
	count := make([]byte, 26)
	for i := 0; i < len(s); i++ {
		delta := s[i] - 'a'
		count[delta]++
	}
	return string(count)
}

// @lc code=end
