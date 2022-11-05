/*
 * @lc app=leetcode id=206 lang=golang
 *
 * [206] Reverse Linked List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	current := head
	var prev, tempNext *ListNode

	for current != nil {
		tempNext = current.Next
		current.Next = prev
		prev = current
		current = tempNext
	}
	return prev
}

// @lc code=end

