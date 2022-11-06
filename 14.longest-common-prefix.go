/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	lenOfStrs := len(strs)
	if lenOfStrs == 0 {
		return ""
	}
	result := strs[0]
	for i := 1; i < lenOfStrs; i++ {
		result = longestPrefix(result, strs[i])
	}

	return result
}

func longestPrefix(str1 string, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)
	var foreachEnd int
	if len1 < len2 {
		foreachEnd = len1
	} else {
		foreachEnd = len2
	}
	endPoint := 0
	for i := 0; i < foreachEnd; i++ {
		if str1[i] != str2[i] {
			break
		}
		endPoint++
	}

	return str1[:endPoint]
}

// @lc code=end

