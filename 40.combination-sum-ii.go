/*
 * @lc app=leetcode id=40 lang=golang
 *
 * [40] Combination Sum II
 */

// @lc code=start
func combinationSum2(candidates []int, target int) [][]int {
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
		if (graph)[i] > target {
			break
		}
		if i > index && (graph)[i] == (graph)[i-1] {
			continue
		}
		dfs(graph, append(path, (graph)[i]), i+1, target-(graph)[i], result)
	}
}

// @lc code=end

