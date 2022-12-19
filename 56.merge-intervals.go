/*
 * @lc app=leetcode id=56 lang=golang
 *
 * [56] Merge Intervals
 */

// @lc code=start
/*
先排序陣列(根據地一個元素大小)
接下來再分三個狀況(取代第二個字 => append一個新的 => 捨去)
*/
func merge(intervals [][]int) [][]int {
	var ans [][]int
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	ans = append(ans, intervals[0])
	for _, v := range intervals[1:] {
		if v[1] >= ans[len(ans)-1][1] {
			if v[0] <= ans[len(ans)-1][1] {
				ans[len(ans)-1][1] = v[1]
			} else {
				ans = append(ans, v)
			}
		}
	}
	return ans
}

// @lc code=end

