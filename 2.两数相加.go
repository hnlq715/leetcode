/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	current := head
	sums := 0
	
	for l1 != nil || l2 != nil || sums>0 {
		if l1 != nil {
			sums += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sums += l2.Val
			l2 = l2.Next
		}

		current.Next = &ListNode{Val: sums%10}
		sums /= 10
		current = current.Next
	}

	return head.Next
}
// @lc code=end

