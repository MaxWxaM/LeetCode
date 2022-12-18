/*
 * @lc app=leetcode id=49 lang=golang
 *
 * [49] Group Anagrams
 */

// @lc code=start
/*
Second version,
1. 開個map去做記錄,
	=>key 是一個array, 裡面包涵26個英文字母出現字數（index = 英文字母, value = count)
	=>value, string的slice
2. 答案就是map的所有value, append 到slice就ok

*/

func groupAnagrams(strs []string) [][]string {
	m := make(map[[26]int][]string)
	for _, str := range strs {
		curr := [26]int{}
		for i := 0; i < len(str); i++ {
			curr[str[i]-'a']++
		}
		m[curr] = append(m[curr], str)
	}
	var result [][]string
	for _, v := range m {
		result = append(result, v)
	}
	return result
}

// @lc code=end

