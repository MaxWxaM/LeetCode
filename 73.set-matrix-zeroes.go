/*
 * @lc app=leetcode id=73 lang=golang
 *
 * [73] Set Matrix Zeroes
 */

// @lc code=start
/*
Second Version(不使用map)
首先遍歷整個矩陣，如果第 i 行為零，則將 matrix[i][0] 標記為 0，如果第 j 列為零，則將 matrix[0][j] 標記為 0。
但有個問題, 第0行或第0列會的標記點會碰撞, 所以需要額外兩個bool標記
*/
func setZeroes(matrix [][]int) {

	if len(matrix) == 0 || len(matrix[0]) == 0 || matrix == nil {
		return
	}
	var zeroRow, zeroCol bool
	m, n := len(matrix), len(matrix[0])

	for row := 0; row < m; row++ {
		for col := 0; col < n; col++ {
			if matrix[row][col] == 0 {
				if row == 0 {
					zeroRow = true
				}
				if col == 0 {
					zeroCol = true
				}
				matrix[row][0] = 0
				matrix[0][col] = 0
			}
		}
	}

	for row := 1; row < m; row++ {
		if matrix[row][0] == 0 {
			for col := 1; col < n; col++ {
				matrix[row][col] = 0
			}
		}
	}

	for col := 1; col < n; col++ {
		if matrix[0][col] == 0 {
			for row := 1; row < m; row++ {
				matrix[row][col] = 0
			}
		}
	}

	if zeroRow == true {
		for col := 1; col < n; col++ {
			matrix[0][col] = 0
		}
	}
	if zeroCol == true {
		for row := 1; row < m; row++ {
			matrix[row][0] = 0
		}
	}
}
// @lc code=end

