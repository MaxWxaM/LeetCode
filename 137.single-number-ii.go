/*
 * @lc app=leetcode id=137 lang=golang
 *
 * [137] Single Number II
 */

// @lc code=start
func singleNumber(nums []int) int {
	sort.Ints(nums)
	l, coun := len(nums), 0
	prev := nums[0]
	ans := prev
	isBreak := false
	for i := 1; i < l; i++ {
		if nums[i] != prev {
			if coun == 0 {
				ans = prev
				isBreak = true
				break
			} else {
				coun = 0
			}
		} else {
			coun++
		}
		prev = nums[i]
	}
	if isBreak == false {
		return prev
	}
	return ans
}

// @lc code=end

