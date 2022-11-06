/*
 * @lc app=leetcode id=1072 lang=golang
 *
 * [1072] Flip Columns For Maximum Number of Equal Rows
 */

// @lc code=start
func maxEqualRowsAfterFlips(matrix [][]int) int {
	lenOfMatrix := len(matrix)
	if lenOfMatrix == 0 {
		return 0
	}
	lenOfMatrixRow := len(matrix[0])
	m := make(map[string]int)
	var isFlip bool
	maxRows := 0
	var key string
	var keyOfByte []byte
	for i := 0; i < lenOfMatrix; i++ {
		if matrix[i][0] == 1 {
			isFlip = true
		} else {
			isFlip = false
		}
		keyOfByte = keyOfByte[:0]
		for j := 1; j < lenOfMatrixRow; j++ {
			keyOfByte = append(keyOfByte, flipByte(matrix[i][j], isFlip))
		}
		key = string(keyOfByte)
		if _, ok := m[key]; ok {
			m[key] += 1
		} else {
			m[key] = 1
		}
		if m[key] > maxRows {
			maxRows = m[key]
		}
	}
	return maxRows
}

func flipByte(i int, isFlip bool) byte {
	if (i == 0 && isFlip) || (i == 1 && !isFlip) {
		return '1'
	}
	return '0'
}

// @lc code=end

