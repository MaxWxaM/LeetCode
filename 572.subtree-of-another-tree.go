/*
 * @lc app=leetcode id=572 lang=golang
 *
 * [572] Subtree of Another Tree
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
 Backtracking
 中斷條件:
  isSubTree會被recusive直到找到root跟subroot的值是一樣的, 接下來再開始比對樹
 root 跟 subRoot都是nil代表順利找到終點 => true
 root跟subRoot只有一個是nil代表樹狀不同 => false

*/
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	return root != nil && (traversal(root, subRoot) || isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot))
}

func traversal(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if root == nil || subRoot == nil {
		return false
	}

	return root.Val == subRoot.Val && traversal(root.Left, subRoot.Left) && traversal(root.Right, subRoot.Right)
}

// @lc code=end

