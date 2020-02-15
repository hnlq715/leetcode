/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	max := 0
	left := 0
	visit := make(map[rune]int)
	for k, v := range s {
		idx, ok := visit[v]
		if ok && left <= idx{
			if k-left > max {
				max = k - left
			}
			left = idx + 1
		}
		visit[v] = k
	}

	if len(s) - left > max {
		max = len(s) - left
	}
	return max
}
// @lc code=end

