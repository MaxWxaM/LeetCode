/*
 * @lc app=leetcode id=139 lang=golang
 *
 * [139] Word Break
 */

// @lc code=start
/*
foundList = s[:i+1] 符合word break
接下來就是看剩下字串是否在dic裏面, 如果沒有就往前一個foundlist去找
*/
func wordBreak(s string, wordDict []string) bool {
	foundList := []int{-1}
	m := map[string]struct{}{}
	for _, v := range wordDict {
		m[v] = struct{}{}
	}
	for i := 0; i < len(s); i++ {
		for j := len(foundList) - 1; j >= 0; j-- {
			if _, ok := m[s[foundList[j]+1:i+1]]; ok {
				foundList = append(foundList, i)
				break
			}
		}
	}
	return foundList[len(foundList)-1] == len(s)-1
}

// @lc code=end

