/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
/*
Init version,
1. 開個map去做記錄, key 是排序過後的字串, value則是slice的位置
2. 用for迴圈遍歷所有str, 先排序過後檢查map中有沒有存在, 已存在就拿map的value去array中append

*/

func groupAnagrams(strs []string) [][]string {
	m := make(map[string]int) // key is
	var result [][]string
	for _, v := range strs {
		sorted := SortString(v)
		if loca, ok := m[sorted]; ok {
			result[loca] = append(result[loca], v)
			continue
		}
		m[sorted] = len(result)
		result = append(result, []string{v})
	}
	return result
}

func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

// @lc code=end

