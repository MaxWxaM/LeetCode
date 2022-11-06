/*
 * @lc app=leetcode id=36 lang=golang
 *
 * [36] Valid Sudoku
 */

// @lc code=start
func isValidSudoku(board [][]byte) bool {
	rowMap := nestedMapStruct{nestedMap: make(map[int]map[byte]struct{})}
	columnMap := nestedMapStruct{nestedMap: make(map[int]map[byte]struct{})}
	smap := squareMapStruct{squareMap: map[square]map[byte]struct{}}
	var b byte
	var squ square
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			b = board[i][j]
			if b == '.' {
				continue
			}
			rowMap.Set(i)
			if _, ok := rowMap.nestedMap[i][b]; ok {
				return false
			}
			rowMap.nestedMap[i][b] = struct{}{}

			columnMap.Set(j)
			if _, ok := columnMap.nestedMap[j][b]; ok {
				return false
			}
			columnMap.nestedMap[j][b] = struct{}{}

			squ = square{x: i / 3, y: j / 3}
			smap.Set(squ)
			if _, ok := smap.squareMap[squ][b]; ok {
				return false
			}
			smap.squareMap[squ][b] = struct{}{}
		}
	}
	return true
}

type square struct {
	x int
	y int
}

type nestedMapStruct struct {
	nestedMap map[int]map[byte]struct{}
}

type squareMapStruct struct {
	squareMap map[square]map[byte]struct{}
}

func (s *nestedMapStruct) Set(i int) {
	child, ok := s.nestedMap[i]
	if !ok {
		child = map[byte]struct{}{}
		s.nestedMap[i] = child
	}
}

func (s *squareMapStruct) Set(sq square) {
	child, ok := s.squareMap[sq]
	if !ok {
		child = map[byte]struct{}{}
		s.squareMap[sq] = child
	}
}

// @lc code=end

