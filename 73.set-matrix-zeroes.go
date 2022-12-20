/*
 * @lc app=leetcode id=73 lang=golang
 *
 * [73] Set Matrix Zeroes
 */

// @lc code=start
/*
1. New 兩個map, 一個存放需要換成0的x軸, 另外一個則是y軸
2. 先遍歷一次新增要置換的資料這兩個map
3. 再遍歷一次去更換
時間複雜度O(2N)
空間複雜度O(2N)
*/
func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	xMap := make(map[int]struct{})
	yMap := make(map[int]struct{})

	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix[0]); y++ {
			if matrix[x][y] == 0 {
				xMap[x] = struct{}{}
				yMap[y] = struct{}{}
			}
		}
	}

	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix[0]); y++ {
			_, xok := xMap[x]
			_, yok := yMap[y]
			if xok || yok {
				matrix[x][y] = 0
			}
		}
	}
}
// @lc code=end

