/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
 */

// @lc code=start
/*
到達某一個點 dp[m][n]的組合數 =  dp [m-1][n] + dp [m][n-1]
意思就是到達他的左邊組合數 + 到達上面的組合數
用此方法就可以透過dp的方法算出答案
需注意的是   rowID or columnID = 0時都只有一種可能
*/
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
				continue
			}
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}
	return dp[m-1][n-1]
}
// @lc code=end

