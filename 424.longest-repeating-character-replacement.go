/*
 * @lc app=leetcode id=424 lang=golang
 *
 * [424] Longest Repeating Character Replacement
 */

// @lc code=start
/*
Sliding window:
假設可取代的char為2
則
4個字裡面至少要2個重複的字
5個字裡面至少要3個重複的字
以此類推可得 right-left-maxCount == k 就是已經爆掉了
e.g. aacbx  => 可取代兩個字   這樣會有一個字不是重複的

假設遇到爆掉的狀況  則left跟right一起往右移
反之則只移動right直到爆掉
*/
func characterReplacement(s string, k int) int {
	maxCount := 0
	ans := 0
	indexCount := [128]int{}
	left, right := 0, 0
	for right < len(s) {
		indexCount[s[right]]++
		maxCount = max(maxCount, indexCount[s[right]])

		if right-left-maxCount == k {
			indexCount[s[left]]--
			left++
		}
		right++
		ans = max(ans, right-left)
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

