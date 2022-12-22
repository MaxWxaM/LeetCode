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
BFS
先廣度搜尋後加到array
再遞迴丟入下一層所有節點
需要ans不斷往下傳
*/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	ans := [][]int{}
	nodes := []*TreeNode{root}
	Traversal(nodes, &ans)
	return ans
}

func Traversal(nodes []*TreeNode, ans *[][]int) {
	arr := make([]int, len(nodes))
	nextNodes := []*TreeNode{}
	for i, v := range nodes {
		arr[i] = v.Val
		if v.Left != nil {
			nextNodes = append(nextNodes, v.Left)
		}
		if v.Right != nil {
			nextNodes = append(nextNodes, v.Right)
		}
	}
	*ans = append(*ans, arr)
	if len(nextNodes) == 0 {
		return
	}

	Traversal(nextNodes, ans)
}

// @lc code=end

