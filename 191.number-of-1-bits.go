/*
 * @lc app=leetcode id=191 lang=golang
 *
 * [191] Number of 1 Bits
 */

// @lc code=start
/*
Bit manipulation
確認第一位是否為1: tempI & 1
做完向右移捨去最右邊的位數
*/

func hammingWeight(num uint32) int {
	totalBits := 0
	for num > 0 {
		if num&1 > 0 {
			totalBits++
		}
		num >>= 1
	}
	return totalBits
}

// @lc code=end

