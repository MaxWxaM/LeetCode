/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 */

// @lc code=start
/*
依序取數字, 在用two pointer取後面兩個, 如果三個加總超過0 則right向左, 反之
多邏輯判斷跳過下個節點為重複的(including left & right, current)
*/
func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums)
	current, left, right := 0, 1, len(nums)-1
	var currentVal int
	for nums[current] <= 0 {
		currentVal = nums[current]
		for left < right {
			if nums[left]+nums[right]+currentVal > 0 {
				for nums[right] == nums[right-1] {
					right--
					if left >= right {
						break
					}
				}
				right--
				continue
			}
			if nums[left]+nums[right]+currentVal < 0 {
				for nums[left] == nums[left+1] {
					left++
					if left >= right {
						break
					}
				}
				left++
				continue
			}
			if nums[left]+nums[right]+currentVal == 0 {
				result = append(result, []int{nums[left], nums[right], currentVal})
				for nums[left] == nums[left+1] {
					left++
					if left >= right {
						break
					}
				}
				left++
				continue
			}
		}
		for nums[current] == nums[current+1] {
			current++
			if current == len(nums)-2 {
				return result
			}
		}
		current++
		if current == len(nums)-2 {
			return result
		}
		left = current + 1
		right = len(nums) - 1
	}
	return result
}

// @lc code=end

