/*
 * @lc app=leetcode id=39 lang=golang
 *
 * [39] Combination Sum
 */

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {
	result := &[][]int{}
	sort.Ints(candidates)
	dfs(candidates, []int{}, 0, target, result)
	return *result
}

func dfs(graph []int, path []int, index int, target int, result *[][]int) {
	if target == 0 {
		tmp := make([]int, len(path))
		copy(tmp, path)
		*result = append(*result, tmp)
		return
	}
	for i := index; i < len(graph); i++ {
		currentVal := graph[i]
		if currentVal > target {
			break
		}
		dfs(graph, append(path, currentVal), i, target-currentVal, result)
	}
}

// @lc code=end

