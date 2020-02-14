/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	pivot := make(map[int]int)
	for i, num := range nums {
		find := target - num
		if idx, ok := pivot[find]; ok {
			return []int{idx,i}
		}
		pivot[num] = i
	}

	return nil
}
// @lc code=end

