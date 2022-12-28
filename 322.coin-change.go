/*
 * @lc app=leetcode id=322 lang=golang
 *
 * [322] Coin Change
 */

// @lc code=start
/*
DP解
開一個memo記錄從1元到預期amount分別要走幾次
現在位置的走幾步基於 現在金額撿掉不同coin去看扣除後的金額走不走的到
走的到的話則用他的步數＋１（要看各個coin扣除後的最小步數)
有個小撇步就是扣除後=0則在dp[0]找到數字0+1
e.g. [2,5], 14
0 : 0(Default Value)
1 : -1
2 : 1 = dp[0] + 1
3 : -1   (dp[3-2] == -1)
4 : 2 = dp[4-2] + 1
5 : 1 = dp[5-5] + 1
6 : 3 = dp[6-2] + 1
7 : 2 = dp[7-2] + 1
8 : 4 = dp[8-2] + 1
9 : 3 = dp[9-2] + 1
10: 2 = dp[10-5] + 1
11: 4 = dp[11-5] + 1
12: 3 = dp[12-5] + 1
13: 5 = dp[13-5] + 1
14: 4 = dp[14-5] + 1
*/
func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)

	for i := 1; i <= amount; i++ {
		min := -1
		for _, c := range coins {
			if c > i || dp[i-c] == -1 {
				continue
			}
			if min == -1 || dp[i-c]+1 < min {
				min = dp[i-c] + 1
				if i == c {
					break
				}
			}
		}
		dp[i] = min
	}
	return dp[amount]
}

// @lc code=end

