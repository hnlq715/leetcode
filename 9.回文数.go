/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	rs := 0
	for i:=x ;i>0; i/=10 {
		rs = rs*10 + i %10
	}

	if rs == x {
		return true
	}

	return false
}
// @lc code=end

