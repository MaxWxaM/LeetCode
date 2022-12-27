/*
 * @lc app=leetcode id=242 lang=golang
 *
 * [242] Valid Anagram
 */

// @lc code=start
/*
DP with memo
1. 開長度為26的陣列, 每個index對應到一個英文字母
2. s的英文字母對應位置出現次數為正數, t為負數
3. 兩者抵銷後,  for loop一次memo應該全部都為0
*/
func isAnagram(s string, t string) bool {
	memo := [26]int{}
	if len(s) != len(t) {
		return false
	}
	for i := 0; i < len(s); i++ {
		memo[s[i]-'a'] += 1
		memo[t[i]-'a'] -= 1
	}

	for _, v := range memo {
		if v != 0 {
			return false
		}
	}
	return true

}

// @lc code=end

