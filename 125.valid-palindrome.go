/*
 * @lc app=leetcode id=125 lang=golang
 *
 * [125] Valid Palindrome
 */

// @lc code=start
/*
two pointer, 唯一的重點就是要記住怎麼轉小寫跟判斷是否為英文...

*/
func isPalindrome(s string) bool {
	for left, right := 0, len(s)-1; left < right; {
		if !isValidChar(s[left]) {
			left++
			continue
		}

		if !isValidChar(s[right]) {
			right--
			continue
		}

		if strings.ToLower(string(s[left])) != strings.ToLower(string(s[right])) {
			return false
		}
		left++
		right--
	}
	return true
}

func isValidChar(ch uint8) bool {
	return (ch >= 65 && ch <= 90) || (ch >= 48 && ch <= 57) || (ch >= 97 && ch <= 122)
}

// @lc code=end

