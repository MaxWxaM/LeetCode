/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 */

// @lc code=start
/*
1. Use DFS to deal with this problem
2. 走過的路要標記起來避免重複走

*/
func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if DFS(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}

func DFS(board [][]byte, word string, i, j int, locationOfString int) bool {
	if i < 0 || j < 0 || i == len(board) || j == len(board[0]) || board[i][j] == '*' {
		return false
	}

	if board[i][j] == word[locationOfString] {
		if locationOfString == len(word)-1 {
			return true
		}
		temp := board[i][j]
		board[i][j] = '*'
		isFound := DFS(board, word, i+1, j, locationOfString+1) || DFS(board, word, i-1, j, locationOfString+1) || DFS(board, word, i, j+1, locationOfString+1) || DFS(board, word, i, j-1, locationOfString+1)
		board[i][j] = temp
		return isFound
	}
	return false
}

// @lc code=end

