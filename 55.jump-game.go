/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
 */

// @lc code=start
func canJump(nums []int) bool {
	maxReachNum := 0
	for i := 0; i < len(nums); i++ {
		if i > maxReachNum {
			return false
		}
		maxReachNum = max(maxReachNum, i+nums[i])
	}
	return true
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// @lc code=end

