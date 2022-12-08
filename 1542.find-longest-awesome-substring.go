/*
 * @lc app=leetcode id=1542 lang=golang
 *
 * [1542] Find Longest Awesome Substring
 */

// @lc code=start
func longestAwesome(s string) int {
	ans, cur, n, m := 0, 0, len(s), map[int]int{0: -1}
	for i := 0; i < n; i++ {
		r := 1 << (strings.IndexByte("0123456789", s[i]) + 1) >> 1
		cur ^= r

		for j := 0; j <= 9; j++ {
			tmp := cur ^ (1 << j)
			if v, ok := m[tmp]; ok {
				ans = max(ans, i-v)
			}
		}

		if _, ok := m[cur]; !ok {
			m[cur] = i
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

