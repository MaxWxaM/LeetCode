/*
 * @lc app=leetcode id=268 lang=golang
 *
 * [268] Missing Number
 */

// @lc code=start
func missingNumber(nums []int) int {
	total, sum := 0, 0
	for i, v := range nums {
		total += i
		sum += v
	}
	total += len(nums)
	return total - sum
}

// @lc code=end

