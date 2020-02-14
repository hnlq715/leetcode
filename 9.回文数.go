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

	len := 0
	for i:=x;i>0;i=i/10 {
		len++
	}

	for idx:=0; idx<(len+1)/2; idx++ {
		fmt.Println(x%(10^idx), x/(10**(len-idx-1)), x,len-idx-1)
		if x/10 == 0 {
			return true
		}

		if x%10 != x/(10**(len-idx-1)) {
			return false
		}

		x = (x%(10**(len-idx-1)))/10
	}

	return true
}
// @lc code=end

