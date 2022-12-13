/*
 * @lc app=leetcode id=2275 lang=golang
 *
 * [2275] Largest Combination With Bitwise AND Greater Than Zero
 */

// @lc code=start
func largestCombination(candidates []int) int {
	maxCount := 0
	m := map[int]int{}
	for _, v := range candidates {
		loc := 0
		for v > 0 {
			if v&1 == 1 {
				m[loc]++
				if m[loc] > maxCount {
					maxCount = m[loc]
				}

			}
			loc++
			v >>= 1
		}
	}
	return maxCount
}

// @lc code=end

