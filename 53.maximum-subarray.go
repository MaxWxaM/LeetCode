/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
 */

// @lc code=start
/*
Init version
1. 只要累積變成負值就直接捨棄變成0
*/

func maxSubArray(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	sum, max := 0, nums[0]

	for _, v := range nums {
		if sum < 0 {
			sum = v
		} else {
			sum += v
		}
		if sum > max {
			max = sum
		}
	}
	return max
}

// @lc code=end

