/*
 * @lc app=leetcode id=217 lang=golang
 *
 * [217] Contains Duplicate
 */

// @lc code=start
/*
//犧牲空間, 開map確認
func containsDuplicate(nums []int) bool {
	m := make(map[int]struct{})
	for _, v := range nums {
		if _, ok := m[v]; ok {
			return true
		}
		m[v] = struct{}{}
	}
	return false
}

*/
//排序=> 時間複雜度O(NlogN), 空間複雜度 O(1)
func containsDuplicate(nums []int) bool {
	sort.Ints(nums)
	prev := nums[0]
	for _, v := range nums[1:] {
		if v == prev {
			return true
		}
		prev = v
	}
	return false
}

// @lc code=end

