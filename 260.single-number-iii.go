/*
 * @lc app=leetcode id=260 lang=golang
 *
 * [260] Single Number III
 */

// @lc code=start
func singleNumber(nums []int) []int {
	sort.Ints(nums)
	l := len(nums)
	ans := make([]int, 0, 2)
	prev, coun := nums[0], 0
	for i := 1; i < l; i++ {
		if nums[i] == prev {
			coun++
		} else {
			if coun == 0 {
				ans = append(ans, prev)
			}
			coun = 0
		}
		prev = nums[i]
	}
	//append last if only one answer
	if len(ans) == 1 {
		ans = append(ans, nums[l-1])
	}

	return ans
}

// @lc code=end

