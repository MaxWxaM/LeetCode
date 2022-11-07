/*
 * @lc app=leetcode id=20 lang=golang
 *
 * [20] Valid Parentheses
 */

// @lc code=start
/*
新增一個slice去存放, 假設有找到一個組合就刪掉
*/
func isValid(s string) bool {
	if len(s) == 1 {
		return false
	} else if len(s) == 0 {
		return true
	}
	m := map[rune]rune{
		'}': '{',
		')': '(',
		']': '[',
	}
	arr := []rune{}
	for _, r := range s {
		switch r {
		case 40, 91, 123:
			arr = append(arr, r)
		case 41, 93, 125:
			if len(arr) == 0 || m[r] != arr[len(arr)-1] {
				return false
			}
			arr = arr[:len(arr)-1]
		}
	}
	return len(arr) == 0
}

// @lc code=end

