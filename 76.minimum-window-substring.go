/*
 * @lc app=leetcode id=76 lang=golang
 *
 * [76] Minimum Window Substring
 */

// @lc code=start
/*
1. 建立need, have := [128]int{}, [128]int{}, 分別去記錄每個字元出現數量
2. 遍歷t將need補齊
3. 透過sliding window => 右邊往右移一格並確認有無找到字元  有的話have++
   如果have=need代表某個char已經找夠了
4. 如果全部字元都找齊, 左指標再跟進直到少一個字元

*/
func minWindow(s string, t string) string {
	need, have := [128]int{}, [128]int{}
	var remainingChar int
	for i := range t {
		if need[t[i]] == 0 {
			remainingChar++
		}
		need[t[i]]++
	}
	targetLen := len(t)
	var minLength = math.MaxInt32
	var ans string
	for left, right := 0, 0; right < len(s); right++ {
		have[s[right]]++
		if need[s[right]] == have[s[right]] {
			remainingChar--
		}

		if remainingChar == 0 {
			for left <= right && left-right+1 <= targetLen && remainingChar == 0 {
				if right-left+1 < minLength {
					minLength = right - left + 1
					ans = s[left : right+1]
				}

				if need[s[left]] == have[s[left]] {
					remainingChar++
				}
				have[s[left]]--
				left++
			}
		}
	}
	return ans
}

// @lc code=end

