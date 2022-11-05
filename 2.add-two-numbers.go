/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
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
	sum := l1.Val + l2.Val
	currentNumber := sum % 10
	carry := sum / 10
	head := &ListNode{Val: currentNumber}
	current := head
	for l1.Next != nil || l2.Next != nil || carry != 0 {
		sum := carry
		if l1.Next != nil {
			sum += l1.Next.Val
			l1 = l1.Next
		}
		if l2.Next != nil {
			sum += l2.Next.Val
			l2 = l2.Next
		}
		currentNumber = sum % 10
		carry = sum / 10
		current.Next = &ListNode{Val: currentNumber}
		current = current.Next
	}
	return head
}

// @lc code=end

