/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
 */

// @lc code=start
func search(nums []int, target int) int {
	l := len(nums)
	left, right := 0, l-1
	var middle int = (left + right) / 2
	if target == nums[left] {
		return left
	}
	if target == nums[right] {
		return right
	}
	for left < right {
		if nums[middle] == target {
			return middle
		}
		if (nums[middle] > nums[left] && (target < nums[middle] && target > nums[left])) ||
			(nums[middle] < nums[left] && (target > nums[left] || target < nums[middle])) {
			right = middle
		} else {
			left = middle
		}
		if left == right-1 {
			return -1
		}
		middle = (left + right) / 2
	}
	return -1
}

// @lc code=end

