/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
/*
Second verion
1. kadane algorithm
*/

func maxSubArray(nums []int) int {
	accu, maxAccu := nums[0], nums[0]
	for _, v := range nums[1:] {
		accu = max(v, accu+v)
		if accu > maxAccu {
			maxAccu = accu
		}
	}
	return maxAccu
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

// @lc code=end

