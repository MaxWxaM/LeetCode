/*
 * @lc app=leetcode id=91 lang=golang
 *
 * [91] Decode Ways
 */

// @lc code=start
/*
DP解, 假設拿到的數字為0, 有幾種可能會導致Parsing失敗
第一個字母為0||前一個字母大>2 e.g. 30|| 前一個字母也為0 e.g. 00
排除這些狀況後, 它會讓組合數變少, 因為會強制跟前一個字母綁定

如果非遇到0, 且與前一個字母組合後<26則是費事數列
如果前一個字母組合>26, 則組合數沒增加
*/
func numDecodings(s string) int {
	count, prevCount := 1, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '0' {
			if i == 0 || s[i-1] > '2' || s[i-1] == '0' {
				return 0
			}
			count, prevCount = prevCount, count
		} else if i > 0 && (s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6')) {
			count, prevCount = count+prevCount, count
		} else {
			prevCount = count
		}
	}
	return count
}

// @lc code=end

