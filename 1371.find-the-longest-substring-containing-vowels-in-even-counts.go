/*
 * @lc app=leetcode id=1371 lang=golang
 *
 * [1371] Find the Longest Substring Containing Vowels in Even Counts
 */

// @lc code=start
func findTheLongestSubstring(s string) int {
	ans, cur, n, m := 0, 0, len(s), map[int]int{0: -1}
	for i := 0; i < n; i++ {
		ib := strings.IndexByte("aeiou", s[i])
		if ib >= 0 {
			r := 1 << (ib + 1) >> 1
			cur ^= r
			if _, ok := m[cur]; !ok {
				m[cur] = i
			}
		}
		ans = max(ans, i-m[cur])
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// @lc code=end

