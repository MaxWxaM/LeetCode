/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	//key : 字元, value : 最後出現位置
	curLen, maxLen, m := 0, 0, make(map[rune]int)
	left := 0//不重複子字串起點
	for i, v := range s {
		if finalLoc, ok := m[v]; ok && finalLoc >= left {
			curLen = i - left // 從前一個字到left的長度
			if curLen > maxLen {
				maxLen = curLen
			}
			left = finalLoc + 1 
		}
		m[v] = i
	}
	curLen = len(s) - left
	if curLen > maxLen {
		maxLen = curLen
	}
	return maxLen
}
// @lc code=end

