/*
 * @lc app=leetcode id=4 lang=golang
 *
 * [4] Median of Two Sorted Arrays
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)

	totalLen := l1 + l2
	isOdd := false
	if totalLen%2 == 0 {
		isOdd = true
	}
	foreachTo := totalLen / 2
	prev := float64(0)
	current := float64(0)
	for i := 0; i <= foreachTo; i++ {
		prev = current
		if l1, l2 := len(nums1), len(nums2); l1 > 0 && l2 > 0 {
			if nums1[0] < nums2[0] {
				current = float64(nums1[0])
				nums1 = nums1[1:]
			} else {
				current = float64(nums2[0])
				nums2 = nums2[1:]
			}
		} else if l1 == 0 {
			current = float64(nums2[0])
			nums2 = nums2[1:]
		} else if l2 == 0 {
			current = float64(nums1[0])
			nums1 = nums1[1:]
		}
	}
	if !isOdd {
		return current
	} else {
		return (current + prev) / float64(2)
	}
}

// @lc code=end

