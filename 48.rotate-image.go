/*
 * @lc app=leetcode id=48 lang=golang
 *
 * [48] Rotate Image
 */

// @lc code=start
func rotate(matrix [][]int) {
	maxLen := len(matrix)

	for i := 0; i <= maxLen-2; i++ {
		for j := 0; j < maxLen-i-1; j++ {
			matrix[i][j], matrix[maxLen-j-1][maxLen-i-1] = matrix[maxLen-j-1][maxLen-i-1], matrix[i][j]
		}
	}

	for i := 0; i < maxLen/2; i++ {
		matrix[i], matrix[maxLen-i-1] = matrix[maxLen-i-1], matrix[i]
	}

}

// @lc code=end

