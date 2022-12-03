/*
 * @lc app=leetcode id=238 lang=golang
 *
 * [238] Product of Array Except Self
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	leftMap := map[int]int{}
	rightMap := map[int]int{}
	lPrev, rPrev := 1, 1
	initL, initR := 0, len(nums)-1
	for l, r := 0, len(nums)-1; initR > l; {
		lCurrent, rCurrent := lPrev*nums[l], rPrev*nums[r]
		leftMap[l] = lCurrent
		rightMap[r] = rCurrent
		lPrev, rPrev = lCurrent, rCurrent
		l++
		r--
	}
	for i := 0; i <= initR; i++ {
		if i == initL {
			result[i] = rightMap[i+1]
		} else if i == initR {
			result[i] = leftMap[i-1]
		} else {
			result[i] = rightMap[i+1] * leftMap[i-1]
		}
	}
	return result
}

// @lc code=end

