/*
 * @lc app=leetcode id=338 lang=golang
 *
 * [338] Counting Bits
 */

// @lc code=start
/*
Bit manipulation
確認第一位是否為1: tempI & 1
做完向右移捨去最右邊的位數
*/
func countBits(n int) []int {
	ans := make([]int, n+1)
	for i := 0; i <= n; i++ {
		totalBits := 0
		tempI := i
		for tempI > 0 {
			totalBits += tempI & 1
			tempI >>= 1
		}
		ans[i] = totalBits
	}
	return ans
}

// @lc code=end

