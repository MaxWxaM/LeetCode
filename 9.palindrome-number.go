/*
 * @lc app=leetcode id=9 lang=golang
 *
 * [9] Palindrome Number
 */

// @lc code=start
func isPalindrome(x int) bool {
	/* Init version - > convert to string and sliding windows to check
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	left, right := 0, len(strconv.Itoa(x))-1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}
	return true
	*/
	/*Return a rotate int and compare*/
	if x < 0 {
		return false
	}
	r := 0
	for z := x; z > 0; z /= 10 {
		tmp := z % 10
		r *= 10
		r += tmp
	}
	return r == x
}

// @lc code=end

