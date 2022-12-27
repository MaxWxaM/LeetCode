/*
 * @lc app=leetcode id=300 lang=golang
 *
 * [300] Longest Increasing Subsequence
 */

// @lc code=start
/*
Patience sorting
https://www.youtube.com/watch?v=l2rCz7skAlk
*/
func lengthOfLIS(nums []int) int {
	patience := []int{nums[0]}

	for _, v := range nums[1:] {
		l := len(patience)
		if v > patience[l-1] {
			patience = append(patience, v)
			continue
		}
		for j := 0; j < l; j++ {
			if v <= patience[j] {
				patience[j] = v
				break
			}
		}
	}
	return len(patience)
}

// @lc code=end

