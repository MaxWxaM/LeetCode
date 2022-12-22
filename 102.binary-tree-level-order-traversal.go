/*
 * @lc app=leetcode id=102 lang=golang
 *
 * [102] Binary Tree Level Order Traversal
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
/*
BFS with QUEUE
*/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ans := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		nodes := queue
		queue = []*TreeNode{}
		arr := make([]int, len(nodes))
		for i, v := range nodes {
			arr[i] = v.Val
			if v.Left != nil {
				queue = append(queue, v.Left)
			}
			if v.Right != nil {
				queue = append(queue, v.Right)
			}
		}
		ans = append(ans, arr)
	}
	return ans
}

// @lc code=end

