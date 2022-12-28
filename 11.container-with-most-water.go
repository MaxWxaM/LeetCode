/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 */

// @lc code=start
/*
Sliding windows
如果右邊高於左邊  則選擇left++
反之right--
因為只有高增加才有可能增加最大水量
*/
func maxArea(height []int) int {
	l := len(height)
	left, right := 0, l-1
	var max int
	for left < right {
		result := min(height[left], height[right]) * (right - left)
		if max < result {
			max = result
		}
		if height[left] > height[right] {
			right--
		} else {
			left++
		}
	}
	return max
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

// @lc code=end

