/*
 * @lc app=leetcode id=153 lang=golang
 *
 * [153] Find Minimum in Rotated Sorted Array
 */

// @lc code=start
/*
Binary search:
最小值在左邊的狀況有兩種:
1 2 3 4 5 => 完全沒轉動, right > mid > left
5 1 2 3 4 => right > mid && left > mid
兩者有共同特徵: right > mid, 所以可推出答案
*/

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right-1 {
		mid := (left + right) / 2
		if nums[mid] < nums[right] {
			right = mid
		} else {
			left = mid
		}
	}
	return min(nums[left], nums[right])
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

// @lc code=end

