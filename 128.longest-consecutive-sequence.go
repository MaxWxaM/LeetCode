/*
 * @lc app=leetcode id=128 lang=golang
 *
 * [128] Longest Consecutive Sequence
 */

// @lc code=start
/*
init version
到某個點候用left, right去左右找 計算count
但效能極差
*/
func longestConsecutive(nums []int) int {
	m := make(map[int]struct{})
	ans := 0
	for _, v := range nums {
		count := 1
		left, right := v-1, v+1
		for {
			if _, ok := m[left]; ok {
				count++
				left--
			} else {
				break
			}
		}
		for {
			if _, ok := m[right]; ok {
				count++
				right++
			} else {
				break
			}
		}
		if count > ans {
			ans = count
		}
		m[v] = struct{}{}
	}
	return ans
}

// @lc code=end

