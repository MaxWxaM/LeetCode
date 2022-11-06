/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/*
預先拉好空間, 在往右走到盡頭 , 此時的left網又跳過一格即是
left                                 right
| ph | -> | 1 | -> | 2 | -> | 3 | -> | 4 | -> | 5 |

	left                             right

| ph | -> | 1 | -> | 2 | -> | 3 | -> | 4 | -> | 5 |   *

	left                    right

| ph | -> | 1 | -> | 2 | -> | 4 | -> | 5 |   *
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	start := &ListNode{Next: head}
	left, right := start, head

	for i := 0; i < n && right != nil; i++ {
		right = right.Next
	}

	for right != nil {
		left, right = left.Next, right.Next
	}

	left.Next = left.Next.Next // skip one node
	return start.Next
}

// @lc code=end

