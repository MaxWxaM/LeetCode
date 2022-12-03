/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */

// @lc code=start
func reverse(x int) int {
	var result int
	result = 0
	for x != 0 {
		result *= 10
		result += x % 10
		x /= 10
	}
	if result > math.MaxInt32 || result < math.MinInt32 {
		return 0
	}
	return result
}

// @lc code=end

