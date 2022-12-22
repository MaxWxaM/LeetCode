/*
 * @lc app=leetcode id=98 lang=golang
 *
 * [98] Validate Binary Search Tree
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
func isValidBST(root *TreeNode) bool {
	min, max := math.MinInt32, math.MaxInt32
	return validBST(root, min, max)

}

/*
DFS:
只要超過允許的設定值就會中斷並回傳false
反之如果都沒有錯誤會走到底到nil

另外left節點的上下線值
在這個case會錯

	    5
		4   6
		   3  7

這個case 的3已經小於跟節點的5, 也就是說要逐層傳遞上下線值
左節點 : min : 上層檢查時的min值,  max:parentnode -1
右節點 : min : parentnode -1,  max:上層檢查時的max值
*/
func validBST(node *TreeNode, min, max int) bool {
	if node == nil {
		return true
	}

	if node.Val < min || node.Val > max {
		return false
	}
	return validBST(node.Left, min, node.Val-1) && validBST(node.Right, node.Val+1, max)
}

// @lc code=end

