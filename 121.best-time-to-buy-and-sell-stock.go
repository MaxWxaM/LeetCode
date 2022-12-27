/*
 * @lc app=leetcode id=121 lang=golang
 *
 * [121] Best Time to Buy and Sell Stock
 */

// @lc code=start
/*
Use two pointer
*/

func maxProfit(prices []int) int {
	min := prices[0]
	max := 0

	for _, v := range prices {
		if v >= min {
			if v-min > max {
				max = v - min
			}
		} else {
			min = v
		}
	}

	return max
}

// @lc code=end

