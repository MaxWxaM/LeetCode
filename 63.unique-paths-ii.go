/*
 * @lc app=leetcode id=63 lang=golang
 *
 * [63] Unique Paths II
 */

// @lc code=start
/*
跟62類似  但多一個條件是一開始進入時就=1代表他有石頭, 這時候就把到達此位置的可能性切成0
另外要多判斷obstacleGrid[0][0] == 1是因為起始就是石頭就直接=0
*/
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	m, n := len(obstacleGrid), len(obstacleGrid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				obstacleGrid[0][0] = 1
				continue
			}
			if obstacleGrid[i][j] == 1 {
				obstacleGrid[i][j] = 0
			} else if i == 0 {
				obstacleGrid[i][j] = obstacleGrid[0][j-1]
			} else if j == 0 {
				obstacleGrid[i][j] = obstacleGrid[i-1][0]
			} else {
				obstacleGrid[i][j] = obstacleGrid[i-1][j] + obstacleGrid[i][j-1]
			}
		}
	}
	return obstacleGrid[m-1][n-1]
}
// @lc code=end

