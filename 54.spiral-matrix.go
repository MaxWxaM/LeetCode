/*
 * @lc app=leetcode id=54 lang=golang
 *
 * [54] Spiral Matrix
 */

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	l, r, t, b := 0, len(matrix[0])-1, 0, len(matrix)-1
	var ans []int
	currentDirection := right
	for l <= r && t <= b {
		switch currentDirection {
		case right:
			for i, j := l, t; i <= r; i++ {
				ans = append(ans, matrix[j][i])
			}
			t++
		case left:
			for i, j := r, b; i >= l; i-- {
				ans = append(ans, matrix[j][i])
			}
			b--
		case up:
			for i, j := l, b; j >= t; j-- {
				ans = append(ans, matrix[j][i])
			}
			l++
		case down:
			for i, j := r, t; j <= b; j++ {
				ans = append(ans, matrix[j][i])
			}
			r--
		}
		currentDirection = changeDirection(currentDirection)
	}
	return ans
}

const (
	right = iota
	down
	left
	up
)

func changeDirection(i int) int {
	if i == up {
		return 0
	}
	return i + 1
}

// @lc code=end

