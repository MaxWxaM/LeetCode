/*
 * @lc app=leetcode id=105 lang=golang
 *
 * [105] Construct Binary Tree from Preorder and Inorder Traversal
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
 遞迴解法:
 1. 首先要思考終止條件:
 假設傳進來的preorder已經是空的即回傳nil
 2. 接下來要把問題divide and conquer
   => 在inorder 中, index < root的即是建立左邊tree所需的inorder
   => 根據inorder取得數量, pre order取第一筆之後的這個數量即是建立左邊tree所需的preorder
   => 右邊同理
 3. 丟下去跑遞迴

 可用 preorder :[1,2,4,5,3,6,8,9,7] 跟 inorder :[4,2,5,1,8,6,9,3,7]下去思考
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	var idx int

	for i, v := range inorder {
		if preorder[0] == v {
			idx = i
			break
		}
	}

	node := &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:idx+1], inorder[:idx]),
		Right: buildTree(preorder[idx+1:], inorder[idx+1:]),
	}
	return node
}

// @lc code=end

