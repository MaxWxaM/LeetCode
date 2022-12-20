/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
 */

// @lc code=start
/*
先找到插入點, 插入後有兩種狀況

1.  if intervals[i][0] > newInterval[1] 無須合併
直接append在中間
2. 從插入點開始向右合併, 遇到終止條件就組裝起來, 另外還有合併到最後一格要考慮, 所以多了這段邏輯j == len(intervals)
*/
func insert(intervals [][]int, newInterval []int) [][]int {
	for i := 0; i < len(intervals); i++ {
		if intervals[i][1] >= newInterval[0] {
			if intervals[i][0] > newInterval[1] {
				result := append(intervals[:i+1], intervals[i:]...)
				result[i] = newInterval
				return result
			}
			for j := i; j <= len(intervals); j++ {
				if j == len(intervals) || intervals[j][0] > newInterval[1] {
					result := append(intervals[:i+1], intervals[j:]...)
					result[i] = newInterval
					return result
				}
				newInterval[0] = min(intervals[j][0], newInterval[0])
				newInterval[1] = max(intervals[j][1], newInterval[1])
			}
		}
	}
	return append(intervals, newInterval)
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}
func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
// @lc code=end

