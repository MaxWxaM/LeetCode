/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	var head, current *ListNode
	if list1.Val < list2.Val {
		head = &ListNode{Val: list1.Val}
		list1 = list1.Next
	} else {
		head = &ListNode{Val: list2.Val}
		list2 = list2.Next
	}
	current = head
	for {
		if list1 == nil {
			current.Next = list2
			break
		} else if list2 == nil {
			current.Next = list1
			break
		} else {
			if list1.Val < list2.Val {
				current.Next = list1
				list1 = list1.Next
			} else {
				current.Next = list2
				list2 = list2.Next
			}
			current = current.Next
		}
	}
	return head
}

// @lc code=end

